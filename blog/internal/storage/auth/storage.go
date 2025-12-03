package auth

import (
	"context"

	proto "github.com/fk4peace/golang_services/shared/grpc_proto/gen"
)

type Storage struct {
	client proto.AuthClient
}

func New(client proto.AuthClient) *Storage {
	return &Storage{client}
}

func (s *Storage) GetPersonRolesById(personId int64) ([]string, error) {
	roles, err := s.client.GetPersonRolesById(context.Background(), &proto.GetPersonRolesByIdRequest{
		PersonId: personId,
	})

	if err != nil {
		return nil, ErrorFrom(err)
	}

	return roles.Roles, nil
}
