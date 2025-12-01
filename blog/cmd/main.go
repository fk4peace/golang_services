package main

import (
	"github.com/fk4peace/golang_services/blog/internal/config"
	"github.com/fk4peace/golang_services/blog/internal/controller/httpController"
	"github.com/fk4peace/golang_services/blog/internal/service"
	"github.com/fk4peace/golang_services/blog/internal/storage/auth"
	"github.com/fk4peace/golang_services/blog/internal/storage/post"
	"github.com/fk4peace/golang_services/blog/pkg/httpServer"
	"github.com/fk4peace/golang_services/blog/pkg/postgresClient"
)

func main() {
	cfg := config.MustLoad()

	dbClient := postgresClient.New(cfg.Postgres)
	postStorage := post.New(dbClient)
	authStorage := auth.New()
	service := service.New(postStorage, authStorage)
	controller := httpController.New(service, cfg.JwtSecret)

	httpServer := httpServer.New(cfg.HttpServer, controller)

	httpServer.Start()
}
