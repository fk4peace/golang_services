package grpcClient

import (
	"fmt"

	"github.com/fk4peace/golang_services/blog/internal/config"
	proto "github.com/fk4peace/golang_services/shared/grpc_proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(cfg config.AuthService) proto.AuthClient {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	return proto.NewAuthClient(conn)
}
