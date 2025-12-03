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

func (e ErrPersonNotFound) Display() string {
	return "person with this id does not exist"
}

type ErrInvalidCredentials struct {
	error
}

func (e ErrInvalidCredentials) Display() string {
	return "invalid username or password"
}

type ErrInvalidRefreshToken struct {
	error
}

func (e ErrInvalidRefreshToken) Display() string {
	return "refresh token is invalid"
}

type ErrUsernameAlreadyTaken struct {
	error
}

func (e ErrUsernameAlreadyTaken) Display() string {
	return "this username is already taken"
}

type ErrInternal struct {
	error
}

func (e ErrInternal) Display() string {
	return "something went wrong, sorry :,("
}
