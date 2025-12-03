package httpController

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (c *httpController) getPostById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		responseError(w, "invalid post id", http.StatusBadRequest)
		return
	}

	post, err := c.service.GetPostById(id)
	if err != nil {
		c.responseErrorFrom(w, err)
		return
	}

	responseData(w, post, http.StatusOK)
}
