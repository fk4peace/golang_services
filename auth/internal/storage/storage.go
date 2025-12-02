package storage

import (
	"context"

	"github.com/fk4peace/golang_services/auth/internal/entity"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	client *pgx.Conn
}

func New(client *pgx.Conn) *Storage {
	return &Storage{
		client,
	}
}

func (s Storage) CreatePerson(username, password string) (*entity.Person, error) {
	query := `
		INSERT INTO person (username, password)
		VALUES ($1, $2)
		RETURNING id, username
	`

	var result entity.Person
	err := s.client.QueryRow(context.Background(), query, username, password).Scan(&result.Id, &result.Username)
	if err != nil {
		return nil, ErrorFrom(err)
	}

	return &result, nil
}

func (s Storage) GetPersonByUsername(username string) (*entity.Person, error) {
	const query = `
		SELECT id, username, password
		FROM person
		WHERE username = $1
	`

	var person entity.Person

	err := s.client.QueryRow(context.Background(), query, username).Scan(
		&person.Id,
		&person.Username,
		&person.Password,
	)
	if err != nil {
		return nil, ErrorFrom(err)
	}

	return &person, nil
}

func (s Storage) GetPersonById(id int64) (*entity.Person, error) {
	const query = `
		SELECT id, username, password
		FROM person
		WHERE id = $1
	`

	var person entity.Person

	err := s.client.QueryRow(context.Background(), query, id).Scan(
		&person.Id,
		&person.Username,
		&person.Password,
	)
	if err != nil {
		return nil, ErrorFrom(err)
	}

	return &person, nil
}

func (s Storage) GetPersonRolesById(personId int64) ([]entity.Role, error) {
	const query = `
		SELECT role
		FROM person
		WHERE id = $1
	`

	var role entity.Role

	err := s.client.QueryRow(context.Background(), query, personId).Scan(&role)
	if err != nil {
		return nil, ErrorFrom(err)
	}

	return []entity.Role{role}, nil
}
