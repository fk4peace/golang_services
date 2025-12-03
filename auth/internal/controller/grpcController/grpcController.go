package grpcController

import (
	"context"

	proto "github.com/fk4peace/golang_services/shared/grpc_proto/gen"
	"go.uber.org/zap"
)

type iService interface {
	GetPersonRolesById(personId int64) ([]string, error)
}

type grpcController struct {
	service iService
	log     *zap.Logger
	proto.UnimplementedAuthServer
}

func New(service iService, log *zap.Logger) *grpcController {
	return &grpcController{
		service: service,
		log:     log,
	}
}

func (c *grpcController) GetPersonRolesById(ctx context.Context, req *proto.GetPersonRolesByIdRequest) (*proto.GetPersonRolesByIdResponse, error) {
	c.log.Info("grpcController", zap.String("endpoint", "GetPersonRolesById"))

	roles, err := c.service.GetPersonRolesById(req.PersonId)
	if err != nil {
		return nil, c.ErrorFrom(err)
	}

	return &proto.GetPersonRolesByIdResponse{Roles: roles}, nil
}
