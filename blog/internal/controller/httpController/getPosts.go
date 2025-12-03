package httpController

import (
	"net/http"
	"strconv"

	"github.com/fk4peace/golang_services/blog/internal/entity"
)

type getPostsResponse struct {
	Total int64
	Posts []entity.Post
}

func (c *httpController) getPosts(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	limit, err := strconv.ParseInt(queryParams.Get("limit"), 10, 64)
	if err != nil {
		limit = 10
	}

	page, err := strconv.ParseInt(queryParams.Get("page"), 10, 64)
	if err != nil {
		page = 1
	}

	posts, total, err := c.service.GetPosts(limit, page)
	if err != nil {
		c.responseErrorFrom(w, err)
		return
	}

	responseData(w, getPostsResponse{
		Total: *total,
		Posts: posts,
	}, http.StatusOK)
}
