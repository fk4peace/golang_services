package service

type ErrInternal struct {
	source error
}

func (e ErrInternal) Error() string {
	return "something went wrong, sorry :,("
}

type ErrPostNotFound struct {
	source error
}

func (e ErrPostNotFound) Error() string {
	return "no posts found with this id"
}

type ErrNoPermissionToPost struct {
	source error
}

func (e ErrNoPermissionToPost) Error() string {
	return "not enough permissions to create post"
}
