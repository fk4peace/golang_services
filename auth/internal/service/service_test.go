package service_test

import (
	"errors"
	"testing"

	"github.com/fk4peace/golang_services/auth/internal/config"
	"github.com/fk4peace/golang_services/auth/internal/entity"
	"github.com/fk4peace/golang_services/auth/internal/service"
	"github.com/fk4peace/golang_services/auth/internal/storage"
	"github.com/stretchr/testify/require"
)

type storageMock struct {
	data    *map[int64]entity.Person
	counter int64
}

func (s storageMock) CreatePerson(username string, password string) (*entity.Person, error) {

	for _, v := range *s.data {
		if v.Username == nil {
			return nil, storage.NewErrInternal(errors.New("error"))
		}

		if *v.Username == username {
			return nil, storage.NewErrAlreadyExists(errors.New("error"))
		}
	}

	id := s.counter
	s.counter++

	(*s.data)[id] = entity.Person{
		Id:       &id,
		Username: &username,
		Password: &password,
	}

	return &entity.Person{
		Id:       &id,
		Username: &username,
	}, nil
}

func (s storageMock) GetPersonById(personId int64) (*entity.Person, error) {
	panic("unimplemented")
}

func (s storageMock) GetPersonByUsername(username string) (*entity.Person, error) {
	panic("unimplemented")
}

func (s storageMock) GetPersonRolesById(personId int64) ([]entity.Role, error) {
	person, ok := (*s.data)[personId]

	if !ok {
		return nil, storage.NewErrNotFound(errors.New("error"))
	}

	if person.Roles == nil {
		return nil, storage.NewErrInternal(errors.New("error"))
	}

	return person.Roles, nil
}

func setup(data *map[int64]entity.Person) *service.Service {
	return service.New(config.Service{
		PasswordSalt:     "egrie783tudsy",
		JwtAccessSecret:  "rvwr3vr3w4",
		JwtRefreshSecret: "kuihj8mc23978qey",
	}, storageMock{data, 1})
}

func TestGetPersonRolesById(t *testing.T) {
	cases := []struct {
		name   string
		data   map[int64]entity.Person
		args   int64
		result []entity.Role
		err    error
	}{
		{
			name: "NoErrors",
			data: map[int64]entity.Person{3: {Roles: []entity.Role{"root", "user"}}},
			args: 3,

			result: []entity.Role{"root", "user"},
		}, {
			name: "ErrorNotFound",
			data: map[int64]entity.Person{},
			args: 5,

			err: service.NewErrPersonNotFound(errors.New("error")),
		}, {
			name: "ErrorInternal",
			data: map[int64]entity.Person{3: {}},
			args: 3,

			err: service.NewErrInternal(errors.New("error")),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			srv := setup(&testCase.data)

			roles, err := srv.GetPersonRolesById(testCase.args)

			if testCase.err == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.err, err)
			}

			if testCase.result == nil {
				require.Nil(t, roles)
			} else {
				require.Equal(t, testCase.result, roles)
			}
		})
	}
}

func TestCreatePerson(t *testing.T) {

	cases := []struct {
		name string
		data map[int64]entity.Person
		args []string

		expectData   map[int64]entity.Person
		expectResult *entity.Person
		expectError  error
	}{
		{
			name: "NoErrors",
			data: map[int64]entity.Person{},
			args: []string{"vasya", "passwd123453"},

			expectData: map[int64]entity.Person{
				1: (func() entity.Person {
					var id int64 = 1
					username := "vasya"
					password := "18cfc5e1aa0bc766932a8a1e3b2009604fa60867782db71321676033f78ed81e"
					return entity.Person{
						Id:       &id,
						Username: &username,
						Password: &password,
					}
				})(),
			},
			expectResult: (func() *entity.Person {
				var id int64 = 1
				username := "vasya"
				return &entity.Person{
					Id:       &id,
					Username: &username,
				}
			})(),
		}, {
			name: "ErrorPasswordTooShort",
			data: map[int64]entity.Person{},
			args: []string{"user", "pass"},

			expectData:  map[int64]entity.Person{},
			expectError: service.ErrPasswordTooShort{},
		}, {
			name: "ErrorUsernameAlreadyTaken",
			data: map[int64]entity.Person{
				1: (func() entity.Person {
					username := "vasya"
					return entity.Person{
						Username: &username,
					}
				})(),
			},
			args: []string{"vasya", "passwd123ederf34w3ew3rew3re"},

			expectData: map[int64]entity.Person{
				1: (func() entity.Person {
					username := "vasya"
					return entity.Person{
						Username: &username,
					}
				})(),
			},
			expectError: service.NewErrUsernameAlreadyTaken(errors.New("error")),
		}, {
			name: "ErrorInternal",
			data: map[int64]entity.Person{1: {}},
			args: []string{"vasya", "passwd123ederf34w3ew3rew3re"},

			expectData:  map[int64]entity.Person{1: {}},
			expectError: service.NewErrInternal(errors.New("error")),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			srv := setup(&testCase.data)

			person, err := srv.CreatePerson(testCase.args[0], testCase.args[1])

			if testCase.expectData != nil {
				require.Equal(t, testCase.expectData, testCase.data)
			}

			if testCase.expectResult == nil {
				require.Nil(t, person)
			} else {
				require.Equal(t, testCase.expectResult, person)
				require.Len(t, testCase.data, 1)
			}

			if testCase.expectError == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.expectError, err)
			}

		})
	}
}
