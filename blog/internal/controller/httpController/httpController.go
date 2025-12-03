package httpController

import (
	"net/http"

	"github.com/fk4peace/golang_services/blog/internal/entity"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
)

type iService interface {
	GetPosts(limit, page int64) ([]entity.Post, *int64, error)
	GetPostById(postId int64) (*entity.Post, error)
	CreatePost(personId int64, content string) (*entity.Post, error)
}

type httpController struct {
	service   iService
	log       zap.Logger
	jwtSecret string
}

func New(service iService, log *zap.Logger, jwtSecret string) http.Handler {
	controller := httpController{
		service,
		*log,
		jwtSecret,
	}

	router := chi.NewRouter()

	router.Use(logger(log))
	router.Get("/posts", controller.getPosts)
	router.Get("/posts/{id}", controller.getPostById)

	router.Group(func(router chi.Router) {
		router.Use(jwtRead(controller.jwtSecret, log))
		router.Post("/posts", controller.createPost)
	})

	return router
}
