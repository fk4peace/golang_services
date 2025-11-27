package service

import (
	"auth_service/internal/storage"
	"errors"
)

type ErrPasswordTooShort struct{}

func (e ErrPasswordTooShort) Error() string {
	return "password must be at least 10 characters"
}

type ErrInvalidCredentials struct{}

func (e ErrInvalidCredentials) Error() string {
	return "invalid username or password"
}

type ErrNotFound struct {
	error
}

type ErrAlreadyExists struct {
	error
}

type ErrInternal struct {
	error
}

func ErrorFrom(err error) error {
	if err == nil {
		return nil
	}

	var errNotFound storage.ErrNotFound
	if errors.As(err, &errNotFound) {
		return ErrNotFound{err}
	}

	var errAlreadyExists storage.ErrAlreadyExists
	if errors.As(err, &errAlreadyExists) {
		return ErrAlreadyExists{err}
	}

	return ErrInternal{err}
}
