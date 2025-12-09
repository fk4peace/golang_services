package storage

type ErrNotFound struct {
	error
}

func NewErrNotFound(err error) ErrNotFound {
	return ErrNotFound{err}
}

type ErrAlreadyExists struct {
	error
}

func NewErrAlreadyExists(err error) ErrAlreadyExists {
	return ErrAlreadyExists{err}
}

type ErrInternal struct {
	error
}

func NewErrInternal(err error) ErrInternal {
	return ErrInternal{err}
}

type ErrUnauthorized struct {
	error
}

func NewErrUnauthorized(err error) ErrUnauthorized {
	return ErrUnauthorized{err}
}

type ErrForbidden struct {
	error
}

func NewErrForbidden(err error) ErrForbidden {
	return ErrForbidden{err}
}

type ErrUnavailable struct {
	error
}

func NewErrUnavailable(err error) ErrUnavailable {
	return ErrUnavailable{err}
}

type ErrBadRequest struct {
	error
}

func NewErrBadRequest(err error) ErrBadRequest {
	return ErrBadRequest{err}
}

type ErrConflict struct {
	error
}

func NewErrConflict(err error) ErrConflict {
	return ErrConflict{err}
}
