package httpController

import (
	"encoding/json"
	"net/http"
)

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type refreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *httpController) refresh(w http.ResponseWriter, r *http.Request) {
	var payload refreshRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		responseErrorInvalidJson(w)
		return
	}

	if payload.RefreshToken == "" {
		responseError(w, "refresh token is required", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := c.service.Refresh(payload.RefreshToken)

	if err != nil {
		responseErrorFrom(w, r, err)
		return
	}

	responseData(w, refreshResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, http.StatusCreated)
}
