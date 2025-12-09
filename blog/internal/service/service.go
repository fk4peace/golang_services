package service

import (
	"errors"
	"slices"

	"github.com/fk4peace/golang_services/blog/internal/entity"
	"github.com/fk4peace/golang_services/blog/internal/storage"
)

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -package=mock -destination=mock/PostsStorage.go -mock_names=iPostsStorage=PostsStorage . iPostsStorage
type iPostsStorage interface {
	GetPosts(limit, page int64) ([]entity.Post, error)
	GetPostsTotal() (*int64, error)
	GetPostById(postId int64) (*entity.Post, error)
	CreatePost(personId int64, content string) (*entity.Post, error)
}

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -package=mock -destination=mock/AuthStorage.go -mock_names=iAuthStorage=AuthStorage . iAuthStorage
type iAuthStorage interface {
	GetPersonRolesById(personId int64) ([]string, error)
}

type Service struct {
	postStore iPostsStorage
	authStore iAuthStorage
}

func New(store iPostsStorage, auth iAuthStorage) *Service {
	return &Service{
		store,
		auth,
	}
}

func (s *Service) GetPosts(limit, page int64) ([]entity.Post, *int64, error) {
	posts, err := s.postStore.GetPosts(limit, page)
	if err != nil {
		return nil, nil, ErrInternal{err}
	}

	total, err := s.postStore.GetPostsTotal()
	if err != nil {
		return nil, nil, ErrInternal{err}
	}

	return posts, total, err
}

func (s *Service) GetPostById(postId int64) (*entity.Post, error) {
	post, err := s.postStore.GetPostById(postId)
	if err != nil {
		var errNotFound storage.ErrNotFound
		if errors.As(err, &errNotFound) {
			return nil, ErrNotPostsFound{err}
		}

		return nil, ErrInternal{err}
	}

	return post, nil
}

func (s *Service) CreatePost(personId int64, content string) (*entity.Post, error) {
	roles, err := s.authStore.GetPersonRolesById(personId)

	if err != nil {
		var errNotFound storage.ErrNotFound
		if errors.As(err, &errNotFound) {
			return nil, ErrNoPersonFound{err}
		}

		return nil, ErrInternal{err}
	}

	if !slices.Contains(roles, "root") {
		return nil, ErrNoPermissionToPost{errors.New("person must have root role")}
	}

	post, err := s.postStore.CreatePost(personId, content)
	if err != nil {
		return nil, ErrInternal{err}
	}

	return post, nil
}
