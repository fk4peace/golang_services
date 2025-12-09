package service

type ErrInternal struct {
	error
}

func NewErrInternal(err error) ErrInternal {
	return ErrInternal{err}
}

func (e ErrInternal) Display() string {
	return "something went wrong, sorry :,("
}

type ErrNotPostsFound struct {
	error
}

func NewErrNotPostsFound(err error) ErrNotPostsFound {
	return ErrNotPostsFound{err}
}

func (e ErrNotPostsFound) Display() string {
	return "no posts found with this id"
}

type ErrNoPersonFound struct {
	error
}

func NewErrNoPersonFound(err error) ErrNoPersonFound {
	return ErrNoPersonFound{err}
}

func (e ErrNoPersonFound) Display() string {
	return "no person found with this id"
}

type ErrNoPermissionToPost struct {
	error
}

func NewErrNoPermissionToPost(err error) ErrNoPermissionToPost {
	return ErrNoPermissionToPost{err}
}

func (e ErrNoPermissionToPost) Display() string {
	return "not enough permissions to create post"
}
