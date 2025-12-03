package grpcController

import (
	"errors"

	"github.com/fk4peace/golang_services/auth/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *grpcController) ErrorFrom(err error) error {
	if err == nil {
		return nil
	}

	c.log.Error("grpcController", zap.Error(err))

	var errPasswordTooShort *service.ErrPasswordTooShort
	if errors.As(err, &errPasswordTooShort) {
		return status.Error(codes.InvalidArgument, errPasswordTooShort.Error())
	}

	var errInvalidCredentials *service.ErrInvalidCredentials
	if errors.As(err, &errInvalidCredentials) {
		return status.Error(codes.Unauthenticated, errInvalidCredentials.Error())
	}

	var errInvalidRefreshToken *service.ErrInvalidRefreshToken
	if errors.As(err, &errInvalidRefreshToken) {
		return status.Error(codes.Unauthenticated, errInvalidRefreshToken.Error())
	}

	var errUsernameAlreadyTaken *service.ErrUsernameAlreadyTaken
	if errors.As(err, &errUsernameAlreadyTaken) {
		return status.Error(codes.AlreadyExists, errUsernameAlreadyTaken.Error())
	}

	var errInternal *service.ErrInternal
	if errors.As(err, &errInternal) {
		return status.Error(codes.Internal, errInternal.Error())
	}

	return status.Error(codes.Internal, "something went wrong, sorry :,(")
}
