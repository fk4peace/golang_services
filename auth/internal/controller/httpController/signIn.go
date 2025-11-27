package httpController

import (
	"auth_service/internal/entity"
	"encoding/json"
	"net/http"
)

type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signInResponse struct {
	Person       entity.Person `json:"person"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}

func (c *httpController) signIn(w http.ResponseWriter, r *http.Request) {
	var payload signInRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		responseErrorInvalidJson(w)
		return
	}

	if payload.Username == "" || payload.Password == "" {
		responseError(w, "username and password are required", http.StatusBadRequest)
		return
	}

	person, err := c.service.SignIn(payload.Username, payload.Password)
	if err != nil {
		responseErrorFrom(w, err)
		return
	}

	accessToken, refreshToken, err := c.service.GenerateTokens(*person.Id)
	if err != nil {
		responseErrorFrom(w, err)
		return
	}

	responseData(w, signInResponse{
		Person:       *person,
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, http.StatusOK)

}
