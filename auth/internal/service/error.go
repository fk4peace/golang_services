package service

type ErrPasswordTooShort struct{}

func (e ErrPasswordTooShort) Error() string {
	return e.Display()
}

func (e ErrPasswordTooShort) Display() string {
	return "password must be at least 10 characters"
}

type ErrPersonNotFound struct {
	error
}

func NewErrPersonNotFound(err error) ErrPersonNotFound {
	return ErrPersonNotFound{err}
}

func (e ErrPersonNotFound) Display() string {
	return "person with this id does not exist"
}

type ErrInvalidCredentials struct {
	error
}

func NewErrInvalidCredentials(err error) ErrInvalidCredentials {
	return ErrInvalidCredentials{err}
}

func (e ErrInvalidCredentials) Display() string {
	return "invalid username or password"
}

type ErrInvalidRefreshToken struct {
	error
}

func NewErrInvalidRefreshToken(err error) ErrInvalidRefreshToken {
	return ErrInvalidRefreshToken{err}
}

func (e ErrInvalidRefreshToken) Display() string {
	return "refresh token is invalid"
}

type ErrUsernameAlreadyTaken struct {
	error
}

func NewErrUsernameAlreadyTaken(err error) ErrUsernameAlreadyTaken {
	return ErrUsernameAlreadyTaken{err}
}

func (e ErrUsernameAlreadyTaken) Display() string {
	return "this username is already taken"
}

type ErrInternal struct {
	error
}

func NewErrInternal(err error) ErrInternal {
	return ErrInternal{err}
}

func (e ErrInternal) Display() string {
	return "something went wrong, sorry :,("
}
