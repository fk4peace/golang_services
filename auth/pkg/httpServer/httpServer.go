package httpServer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fk4peace/golang_services/auth/internal/config"
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
	log.Println("HttpServer running on :" + strconv.Itoa(server.port))

	err := http.ListenAndServe(":"+strconv.Itoa(server.port), *server.handler)

	if err != nil {
		panic(err)
	}
}
