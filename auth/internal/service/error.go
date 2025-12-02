package service

type ErrPasswordTooShort struct{}

func (e ErrPasswordTooShort) Error() string {
	return "password must be at least 10 characters"
}

type ErrPersonNotFound struct {
	source error
}

func (e ErrPersonNotFound) Error() string {
	return "person with this id does not exist"
}

type ErrInvalidCredentials struct {
	source error
}

func (e ErrInvalidCredentials) Error() string {
	return "invalid username or password"
}

type ErrInvalidRefreshToken struct {
	source error
}

func (e ErrInvalidRefreshToken) Error() string {
	return "refresh token is invalid"
}

type ErrUsernameAlreadyTaken struct {
	source error
}

func (e ErrUsernameAlreadyTaken) Error() string {
	return "this username is already taken"
}

type ErrInternal struct {
	source error
}

func (e ErrInternal) Error() string {
	return "something went wrong, sorry :,("
}
