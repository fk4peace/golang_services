package httpController

import (
	"net/http"

	"github.com/fk4peace/golang_services/auth/internal/entity"

	"github.com/go-chi/chi/v5"
)

type iservice interface {
	CreatePerson(username, password string) (*entity.Person, error)
	GenerateTokens(id int64) (*string, *string, error)
	SignIn(username, password string) (*entity.Person, error)
	Refresh(refreshToken string) (*string, *string, error)
}

type httpController struct {
	service iservice
}

func New(service iservice) http.Handler {
	controller := httpController{
		service: service,
	}

	router := chi.NewRouter()

	router.Post("/sign_up", controller.signUp)
	router.Post("/sign_in", controller.signIn)
	router.Post("/refresh", controller.refresh)

	return router
}
