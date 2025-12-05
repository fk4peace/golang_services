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
		return &storage.ErrUnavailable{Source: err}
	}

	switch st.Code() {
	case codes.NotFound:
		return &storage.ErrNotFound{Source: err}

	case codes.InvalidArgument:
		return &storage.ErrBadRequest{Source: err}

	case codes.AlreadyExists:
		return &storage.ErrConflict{Source: err}

	case codes.Unauthenticated:
		return &storage.ErrUnauthorized{Source: err}

	case codes.PermissionDenied:
		return &storage.ErrForbidden{Source: err}

	case codes.Unavailable:
		return &storage.ErrUnavailable{Source: err}

	case codes.DeadlineExceeded:
		return &storage.ErrUnavailable{Source: err}

	case codes.Internal:
		return &storage.ErrInternal{Source: err}

	default:
		return &storage.ErrInternal{Source: err}
	}
}
