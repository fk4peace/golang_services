package post

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

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

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound{err}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return ErrAlreadyExists{err}
	}

	return ErrInternal{err}
}
