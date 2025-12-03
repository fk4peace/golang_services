package main

import (
	"github.com/fk4peace/golang_services/auth/pkg/logger"
	"github.com/fk4peace/golang_services/blog/internal/config"
	"github.com/fk4peace/golang_services/blog/internal/controller/httpController"
	"github.com/fk4peace/golang_services/blog/internal/service"
	"github.com/fk4peace/golang_services/blog/internal/storage/auth"
	"github.com/fk4peace/golang_services/blog/internal/storage/post"
	"github.com/fk4peace/golang_services/blog/pkg/grpcClient"
	"github.com/fk4peace/golang_services/blog/pkg/httpServer"
	"github.com/fk4peace/golang_services/blog/pkg/postgresClient"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.New()

	dbClient := postgresClient.New(cfg.Postgres)
	grpcClient := grpcClient.New(cfg.AuthService)

	postStorage := post.New(dbClient)
	authStorage := auth.New(grpcClient)

	service := service.New(postStorage, authStorage)

	controller := httpController.New(service, logger, cfg.JwtSecret)

	httpServer := httpServer.New(cfg.HttpServer, controller)

	httpServer.Start()
}
