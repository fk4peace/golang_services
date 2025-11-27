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
	var errPasswordTooShort service.ErrPasswordTooShort
	if errors.As(err, &errPasswordTooShort) {
		responseError(writer, err.Error(), 422)
		return
	}

	var errInvalidCredentials service.ErrInvalidCredentials
	if errors.As(err, &errInvalidCredentials) {
		responseError(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var errInvalidRefreshToken service.ErrInvalidRefreshToken
	if errors.As(err, &errInvalidRefreshToken) {
		responseError(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var errAlreadyExists service.ErrUsernameAlreadyTaken
	if errors.As(err, &errAlreadyExists) {
		responseError(writer, err.Error(), http.StatusConflict)
		return
	}

	var ErrInternal service.ErrInternal
	if errors.As(err, &ErrInternal) {
		responseError(writer, err.Error(), 500)
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
