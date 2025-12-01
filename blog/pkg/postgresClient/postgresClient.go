package postgresClient

import (
	"context"
	"fmt"

	"github.com/fk4peace/golang_services/blog/internal/config"

	"github.com/jackc/pgx/v5"
)

func New(cfg config.Postgres) *pgx.Conn {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		panic((err))
	}

	err = conn.Ping(ctx)
	if err != nil {
		panic((err))
	}

	return conn
}
