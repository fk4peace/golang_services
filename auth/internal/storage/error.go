package storage

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ErrNotFound struct {
	error
}

func NewErrNotFound(err error) ErrNotFound {
	return ErrNotFound{err}
}

type ErrAlreadyExists struct {
	error
}

func NewErrAlreadyExists(err error) ErrAlreadyExists {
	return ErrAlreadyExists{err}
}

type ErrInternal struct {
	error
}

func NewErrInternal(err error) ErrInternal {
	return ErrInternal{err}
}

func ErrorFrom(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound{err}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return ErrAlreadyExists{err}
	}

	return ErrInternal{err}
}
