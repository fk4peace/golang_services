package httpController

import (
	"net/http"

	"github.com/fk4peace/golang_services/auth/internal/entity"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
)

type iService interface {
	CreatePerson(username, password string) (*entity.Person, error)
	GenerateTokens(personId int64) (*string, *string, error)
	SignIn(username, password string) (*entity.Person, error)
	Refresh(refreshToken string) (*string, *string, error)
}

type httpController struct {
	service iService
}

func New(service iService, log *zap.Logger) http.Handler {
	controller := httpController{
		service: service,
	}

	router := chi.NewRouter()

	router.Use(logger(log))

	router.Post("/sign_up", controller.signUp)
	router.Post("/sign_in", controller.signIn)
	router.Post("/refresh", controller.refresh)

	return router
}
