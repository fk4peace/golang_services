package main

import (
	"auth_service/internal/config"
	"auth_service/internal/controller/httpController"
	"auth_service/internal/service"
	"auth_service/internal/storage"
	"auth_service/lib/httpServer"
	"auth_service/lib/postgresClient"
)

func main() {
	cfg := config.MustLoad()

	dbClient := postgresClient.New(cfg.Postgres)
	storage := storage.New(dbClient)
	service := service.New(cfg.Service, storage)
	controller := httpController.New(service)

	httpServer := httpServer.New(cfg.HttpServer, controller)

	httpServer.Start()
}
