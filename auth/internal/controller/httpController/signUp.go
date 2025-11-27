package httpController

import (
	"auth_service/internal/entity"
	"auth_service/internal/service"
	"encoding/json"
	"errors"
	"net/http"
)

type signUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signUpResponse struct {
	Person       entity.Person `json:"person"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}

func (c *httpController) signUp(w http.ResponseWriter, r *http.Request) {
	var payload signUpRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		responseErrorInvalidJson(w)
		return
	}

	if payload.Username == "" || payload.Password == "" {
		responseError(w, "username and password are required", http.StatusBadRequest)
		return
	}

	person, err := c.service.CreatePerson(payload.Username, payload.Password)

	if err != nil {
		var errAlreadyExists service.ErrAlreadyExists
		if errors.As(err, &errAlreadyExists) {
			responseError(w, "person with this username already exists", http.StatusConflict)
			return
		}

		responseErrorFrom(w, err)
		return
	}

	accessToken, refreshToken, err := c.service.GenerateTokens(*person.Id)

	if err != nil {
		responseErrorFrom(w, err)
		return
	}

	responseData(w, signUpResponse{
		Person:       *person,
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, http.StatusCreated)
}
