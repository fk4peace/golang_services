package httpController

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fk4peace/golang_services/blog/internal/service"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func responseErrorFrom(writer http.ResponseWriter, err error) {
	var ErrPostNotFound service.ErrPostNotFound
	if errors.As(err, &ErrPostNotFound) {
		responseError(writer, err.Error(), http.StatusNotFound)
		return
	}

	var ErrInternal service.ErrInternal
	if errors.As(err, &ErrInternal) {
		responseError(writer, err.Error(), http.StatusInternalServerError)
		return
	}

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
