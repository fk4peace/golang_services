package grpcController

import (
	"context"

	proto "github.com/fk4peace/golang_services/shared/grpc_proto/gen"
)

type iService interface {
	GetPersonRolesById(personId int64) ([]string, error)
}

type grpcController struct {
	service iService
	proto.UnimplementedAuthServer
}

func New(service iService) *grpcController {
	return &grpcController{service: service}
}

func (c *grpcController) GetPersonRolesById(ctx context.Context, req *proto.GetPersonRolesByIdRequest) (*proto.GetPersonRolesByIdResponse, error) {
	roles, err := c.service.GetPersonRolesById(req.PersonId)
	if err != nil {
		return nil, err
	}

	return &proto.GetPersonRolesByIdResponse{Roles: roles}, nil
}
