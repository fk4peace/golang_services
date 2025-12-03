package service

type ErrInternal struct {
	error
}

func (e ErrInternal) Display() string {
	return "something went wrong, sorry :,("
}

type ErrNotFound struct {
	error
}

func (e ErrNotFound) Display() string {
	return "no posts found with this id"
}

type ErrNoPermissionToPost struct {
	error
}

func (e ErrNoPermissionToPost) Display() string {
	return "not enough permissions to create post"
}
