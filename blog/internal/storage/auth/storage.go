package auth

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) GetPersonRolesById(personId int64) ([]string, error) {
	return []string{"s"}, nil
}
