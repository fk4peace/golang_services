package main

import (
	"github.com/fk4peace/golang_services/auth/internal/config"
	"github.com/fk4peace/golang_services/auth/internal/controller/grpcController"
	"github.com/fk4peace/golang_services/auth/internal/controller/httpController"
	"github.com/fk4peace/golang_services/auth/internal/service"
	"github.com/fk4peace/golang_services/auth/internal/storage"
	"github.com/fk4peace/golang_services/auth/pkg/grpcServer"
	"github.com/fk4peace/golang_services/auth/pkg/httpServer"
	"github.com/fk4peace/golang_services/auth/pkg/postgresClient"
)

func main() {
	cfg := config.MustLoad()

	dbClient := postgresClient.New(cfg.Postgres)
	storage := storage.New(dbClient)
	service := service.New(cfg.Service, storage)
	httpController := httpController.New(service)
	grpcController := grpcController.New(service)

	httpServer := httpServer.New(cfg.HttpServer, httpController)
	grpcServer := grpcServer.New(cfg.GrpcServer, grpcController)

	go httpServer.Start()

	grpcServer.Start()
}
