package httpController

import (
	"encoding/json"
	"net/http"
)

type createPostRequest struct {
	Content string `json:"content"`
}

func (c *httpController) createPost(w http.ResponseWriter, r *http.Request) {
	personId := requestPerson(r)
	if personId == nil {
		responseError(w, "something went wrong, sorry :,(", http.StatusInternalServerError)
		return
	}

	var payload createPostRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		responseErrorInvalidJson(w)
		return
	}

	if payload.Content == "" {
		responseError(w, "post content is required", http.StatusBadRequest)
		return
	}

	post, err := c.service.CreatePost(*personId, payload.Content)
	if err != nil {
		responseErrorFrom(w, err)
		return
	}

	responseData(w, post, http.StatusOK)
}
