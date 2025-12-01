package service

import (
	"errors"

	"github.com/fk4peace/golang_services/blog/internal/entity"
	postStorage "github.com/fk4peace/golang_services/blog/internal/storage/post"
)

type iPostsStorage interface {
	GetPosts(limit, page int64) ([]entity.Post, error)
	GetPostsTotal() (*int64, error)
	GetPostById(postId int64) (*entity.Post, error)
	CreatePost(personId int64, content string) (*entity.Post, error)
}

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
		var ErrNotFound postStorage.ErrNotFound
		if errors.As(err, &ErrNotFound) {
			return nil, ErrPostNotFound{err}
		}

		return nil, ErrInternal{err}
	}

	return post, nil
}

func (s *Service) CreatePost(personId int64, content string) (*entity.Post, error) {
	post, err := s.postStore.CreatePost(personId, content)
	if err != nil {
		return nil, ErrInternal{err}
	}

	return post, nil
}
