package storage

type ErrNotFound struct {
	Source error
}

func (e ErrNotFound) Error() string {
	return e.Source.Error()
}

type ErrAlreadyExists struct {
	Source error
}

func (e ErrAlreadyExists) Error() string {
	return e.Source.Error()
}

type ErrInternal struct {
	Source error
}

func (e ErrInternal) Error() string {
	return e.Source.Error()
}

type ErrUnauthorized struct {
	Source error
}

func (e ErrUnauthorized) Error() string {
	return e.Source.Error()
}

type ErrForbidden struct {
	Source error
}

func (e ErrForbidden) Error() string {
	return e.Source.Error()
}

type ErrUnavailable struct {
	Source error
}

func (e ErrUnavailable) Error() string {
	return e.Source.Error()
}

type ErrBadRequest struct {
	Source error
}

func (e ErrBadRequest) Error() string {
	return e.Source.Error()
}

type ErrConflict struct {
	Source error
}

func (e ErrConflict) Error() string {
	return e.Source.Error()
}
