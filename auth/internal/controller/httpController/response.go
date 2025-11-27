package httpController

import (
	"auth_service/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func responseErrorFrom(writer http.ResponseWriter, err error) {
	var ErrNotFound service.ErrNotFound
	if errors.As(err, &ErrNotFound) {
		responseError(writer, "person not found", 404)
		return
	}

	var ErrPasswordTooShort service.ErrPasswordTooShort
	if errors.As(err, &ErrPasswordTooShort) {
		responseError(writer, err.Error(), 422)
		return
	}

	var ErrInvalidCredentials service.ErrInvalidCredentials
	if errors.As(err, &ErrInvalidCredentials) {
		responseError(writer, err.Error(), 401)
		return
	}

	var ErrInternal service.ErrInternal
	if errors.As(err, &ErrInternal) {
		responseError(writer, "something went wrong, sorry :,(", 500)
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

	writer.Write(result)
	writer.WriteHeader(status)
}
