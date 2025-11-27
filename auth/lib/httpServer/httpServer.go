package httpServer

import (
	"auth_service/internal/config"
	"net/http"
	"strconv"
)

type HttpServer struct {
	port    int
	handler *http.Handler
}

func New(cfg config.HttpServer, handler http.Handler) *HttpServer {
	return &HttpServer{
		port:    cfg.Port,
		handler: &handler,
	}
}

func (server *HttpServer) Start() {
	err := http.ListenAndServe(":"+strconv.Itoa(server.port), *server.handler)

	if err != nil {
		panic(err)
	}
}
