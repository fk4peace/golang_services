package service_test

import (
	"errors"
	"testing"

	"github.com/fk4peace/golang_services/blog/internal/entity"
	"github.com/fk4peace/golang_services/blog/internal/service"
	"github.com/fk4peace/golang_services/blog/internal/service/mock"
	"github.com/fk4peace/golang_services/blog/internal/storage"
	"github.com/fk4peace/golang_services/blog/pkg"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetPosts(t *testing.T) {

	cases := []struct {
		name  string
		args  []int64
		setup func(*mock.PostsStorage, *mock.AuthStorage)

		expectResultPosts []entity.Post
		expectResultTotal *int64
		expectError       error
	}{
		{
			name: "NoErrors",
			args: []int64{3, 0},
			setup: func(posts *mock.PostsStorage, _ *mock.AuthStorage) {
				posts.EXPECT().
					GetPosts(gomock.Any(), gomock.Any()).
					Return([]entity.Post{{}, {}, {}}, nil).
					AnyTimes()

				posts.EXPECT().
					GetPostsTotal().
					Return(pkg.Ptr(int64(10)), nil).
					AnyTimes()
			},

			expectResultPosts: []entity.Post{{}, {}, {}},
			expectResultTotal: pkg.Ptr(int64(10)),
		},
		{
			name: "ErrorGetPostsInternal",
			args: []int64{3, 0},
			setup: func(posts *mock.PostsStorage, _ *mock.AuthStorage) {
				posts.EXPECT().
					GetPosts(gomock.Any(), gomock.Any()).
					Return(nil, storage.NewErrBadRequest(errors.New("error"))).
					AnyTimes()

				posts.EXPECT().
					GetPostsTotal().
					Return(pkg.Ptr(int64(10)), nil).
					AnyTimes()
			},

			expectError: service.NewErrInternal(errors.New("error")),
		},
		{
			name: "ErrorGetPostsTotalInternal",
			args: []int64{3, 0},
			setup: func(posts *mock.PostsStorage, _ *mock.AuthStorage) {
				posts.EXPECT().
					GetPosts(gomock.Any(), gomock.Any()).
					Return([]entity.Post{{}, {}, {}}, nil).
					AnyTimes()

				posts.EXPECT().
					GetPostsTotal().
					Return(nil, storage.NewErrBadRequest(errors.New("error"))).
					AnyTimes()
			},

			expectError: service.NewErrInternal(errors.New("error")),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			postsStorageMock := mock.NewPostsStorage(ctrl)
			authStorageMock := mock.NewAuthStorage(ctrl)

			testCase.setup(postsStorageMock, authStorageMock)

			srv := service.New(
				postsStorageMock,
				authStorageMock,
			)

			posts, total, err := srv.GetPosts(testCase.args[0], testCase.args[1])

			if testCase.expectError == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.expectError, err)
			}

			if testCase.expectResultPosts == nil {
				require.Nil(t, posts)
			} else {
				require.Equal(t, testCase.expectResultPosts, posts)
			}

			if testCase.expectResultTotal == nil {
				require.Nil(t, total)
			} else {
				require.Equal(t, testCase.expectResultTotal, total)
			}

			ctrl.Finish()

		})
	}

}

func TestGetPostById(t *testing.T) {

	cases := []struct {
		name  string
		args  int64
		setup func(*mock.PostsStorage, *mock.AuthStorage)

		expectResult *entity.Post
		expectError  error
	}{
		{
			name: "NoErrors",
			args: int64(3),
			setup: func(posts *mock.PostsStorage, _ *mock.AuthStorage) {
				posts.EXPECT().
					GetPostById(gomock.Any()).
					Return(&entity.Post{Id: pkg.Ptr(int64(3))}, nil).
					AnyTimes()
			},

			expectResult: &entity.Post{Id: pkg.Ptr(int64(3))},
		},
		{
			name: "ErrorNotFound",
			args: int64(3),
			setup: func(posts *mock.PostsStorage, _ *mock.AuthStorage) {
				posts.EXPECT().
					GetPostById(gomock.Any()).
					Return(nil, storage.NewErrNotFound(errors.New("error"))).
					AnyTimes()
			},

			expectError: service.NewErrNotPostsFound(errors.New("error")),
		},
		{
			name: "ErrorInternal",
			args: int64(3),
			setup: func(posts *mock.PostsStorage, _ *mock.AuthStorage) {
				posts.EXPECT().
					GetPostById(gomock.Any()).
					Return(nil, storage.NewErrInternal(errors.New("error"))).
					AnyTimes()
			},

			expectError: service.NewErrInternal(errors.New("error")),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			postsStorageMock := mock.NewPostsStorage(ctrl)
			authStorageMock := mock.NewAuthStorage(ctrl)

			testCase.setup(postsStorageMock, authStorageMock)

			srv := service.New(
				postsStorageMock,
				authStorageMock,
			)

			post, err := srv.GetPostById(testCase.args)

			if testCase.expectError == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.expectError, err)
			}

			if testCase.expectResult == nil {
				require.Nil(t, post)
			} else {
				require.Equal(t, testCase.expectResult, post)
			}

			ctrl.Finish()
		})
	}

}

func TestCreatePost(t *testing.T) {

	cases := []struct {
		name  string
		arg1  int64
		arg2  string
		setup func(*mock.PostsStorage, *mock.AuthStorage)

		expectResult *entity.Post
		expectError  error
	}{
		{
			name: "NoErrors",
			arg1: 5,
			arg2: "content",
			setup: func(posts *mock.PostsStorage, auth *mock.AuthStorage) {
				auth.EXPECT().
					GetPersonRolesById(gomock.Any()).
					Return([]string{"root"}, nil).
					AnyTimes()

				posts.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Return(&entity.Post{
						Id:       pkg.Ptr(int64(3)),
						PersonId: pkg.Ptr(int64(5)),
						Content:  pkg.Ptr("content"),
					}, nil).
					AnyTimes()
			},

			expectResult: &entity.Post{
				Id:       pkg.Ptr(int64(3)),
				PersonId: pkg.Ptr(int64(5)),
				Content:  pkg.Ptr("content"),
			},
		},
		{
			name: "ErrorPersonNotFound",
			arg1: 5,
			arg2: "content",
			setup: func(posts *mock.PostsStorage, auth *mock.AuthStorage) {
				auth.EXPECT().
					GetPersonRolesById(gomock.Any()).
					Return(nil, storage.NewErrNotFound(errors.New("errors"))).
					AnyTimes()
			},

			expectError: service.NewErrNoPersonFound(errors.New("errors")),
		},
		{
			name: "GetPersonErrorInternal",
			arg1: 5,
			arg2: "content",
			setup: func(posts *mock.PostsStorage, auth *mock.AuthStorage) {
				auth.EXPECT().
					GetPersonRolesById(gomock.Any()).
					Return(nil, storage.NewErrBadRequest(errors.New("errors"))).
					AnyTimes()
			},

			expectError: service.NewErrInternal(errors.New("errors")),
		},
		{
			name: "ErrorNoPermissionToPost",
			arg1: 5,
			arg2: "content",
			setup: func(posts *mock.PostsStorage, auth *mock.AuthStorage) {
				auth.EXPECT().
					GetPersonRolesById(gomock.Any()).
					Return([]string{}, nil).
					AnyTimes()
			},

			expectError: service.NewErrNoPermissionToPost(errors.New("errors")),
		},
		{
			name: "CreatePostErrorInternal",
			arg1: 5,
			arg2: "content",
			setup: func(posts *mock.PostsStorage, auth *mock.AuthStorage) {
				auth.EXPECT().
					GetPersonRolesById(gomock.Any()).
					Return([]string{"root"}, nil).
					AnyTimes()

				posts.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Return(nil, storage.NewErrBadRequest(errors.New("error"))).
					AnyTimes()
			},

			expectError: service.NewErrInternal(errors.New("errors")),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			postsStorageMock := mock.NewPostsStorage(ctrl)
			authStorageMock := mock.NewAuthStorage(ctrl)

			testCase.setup(postsStorageMock, authStorageMock)

			srv := service.New(
				postsStorageMock,
				authStorageMock,
			)

			post, err := srv.CreatePost(testCase.arg1, testCase.arg2)

			if testCase.expectError == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.expectError, err)
			}

			if testCase.expectResult == nil {
				require.Nil(t, post)
			} else {
				require.Equal(t, testCase.expectResult, post)
			}

			ctrl.Finish()
		})
	}

}
