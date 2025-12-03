package post

import (
	"errors"

	"github.com/fk4peace/golang_services/blog/internal/storage"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func ErrorFrom(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return storage.ErrNotFound{Source: err}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return storage.ErrAlreadyExists{Source: err}
	}

	return storage.ErrInternal{Source: err}
}
