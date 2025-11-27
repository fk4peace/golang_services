package httpController

import (
	"auth_service/internal/entity"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type iservice interface {
	CreatePerson(username, password string) (*entity.Person, error)
	GenerateTokens(id int64) (*string, *string, error)
	SignIn(username, password string) (*entity.Person, error)
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

	return router
}
