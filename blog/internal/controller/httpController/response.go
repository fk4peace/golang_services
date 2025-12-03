package httpController

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fk4peace/golang_services/blog/internal/service"
	"go.uber.org/zap"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func (c *httpController) responseErrorFrom(writer http.ResponseWriter, err error) {
	c.log.Error("httpController", zap.Error(err))

	var ErrNotFound service.ErrNotFound
	if errors.As(err, &ErrNotFound) {
		responseError(writer, ErrNotFound.Display(), http.StatusNotFound)
		return
	}

	var errNoPermissionToPost service.ErrNoPermissionToPost
	if errors.As(err, &errNoPermissionToPost) {
		responseError(writer, errNoPermissionToPost.Display(), http.StatusForbidden)
		return
	}

	responseError(writer, "something went wrong, sorry :,(", http.StatusInternalServerError)
}

func responseErrorInvalidJson(writer http.ResponseWriter) {
	responseError(writer, "invalid JSON", http.StatusBadRequest)
}

func responseError(writer http.ResponseWriter, message string, status int) {

	if status < 400 {
		panic(fmt.Sprintf("the http code %d is not an error code", status))
	}

	result, err := json.Marshal(Response{
		Error: message,
	})

	if err != nil {
		panic(err)
	}

	writer.WriteHeader(status)
	writer.Write(result)
}

func responseData(writer http.ResponseWriter, data interface{}, status int) {

	if status < 200 || status > 299 {
		panic(fmt.Sprintf("http code %d is not a success code", status))
	}

	result, err := json.Marshal(Response{
		Data: data,
	})

	if err != nil {
		panic(err)
	}

	writer.WriteHeader(status)
	writer.Write(result)
}
