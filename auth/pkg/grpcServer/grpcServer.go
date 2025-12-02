package grpcServer

import (
	"log"
	"net"
	"strconv"

	"github.com/fk4peace/golang_services/auth/internal/config"
	proto "github.com/fk4peace/golang_services/shared/grpc_proto/gen"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	port    int
	handler proto.AuthServer
}

func New(cfg config.GrpcServer, handler proto.AuthServer) *GrpcServer {
	return &GrpcServer{
		port:    cfg.Port,
		handler: handler,
	}
}

func (server *GrpcServer) Start() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAuthServer(grpcServer, server.handler)

	log.Println("GrpcServer running on :" + strconv.Itoa(server.port))
	grpcServer.Serve(lis)
}
