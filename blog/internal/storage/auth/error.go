package auth

import (
	"github.com/fk4peace/golang_services/blog/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorFrom(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return storage.NewErrUnavailable(err)
	}

	switch st.Code() {
	case codes.NotFound:
		return storage.NewErrNotFound(err)

	case codes.InvalidArgument:
		return storage.NewErrBadRequest(err)

	case codes.AlreadyExists:
		return storage.NewErrConflict(err)

	case codes.Unauthenticated:
		return storage.NewErrUnauthorized(err)

	case codes.PermissionDenied:
		return storage.NewErrForbidden(err)

	case codes.Unavailable:
		return storage.NewErrUnavailable(err)

	case codes.DeadlineExceeded:
		return storage.NewErrUnavailable(err)

	case codes.Internal:
		return storage.NewErrInternal(err)

	default:
		return storage.NewErrInternal(err)
	}
}
