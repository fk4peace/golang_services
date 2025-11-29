package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/fk4peace/golang_services/auth/internal/config"
	"github.com/fk4peace/golang_services/auth/internal/entity"
	"github.com/fk4peace/golang_services/auth/internal/storage"

	"github.com/golang-jwt/jwt/v5"
)

type istorage interface {
	CreatePerson(username, password string) (*entity.Person, error)
	GetPersonByUsername(username string) (*entity.Person, error)
	GetPersonById(id int64) (*entity.Person, error)
}

type Service struct {
	store  istorage
	config config.Service
}

func New(cfg config.Service, storage istorage) *Service {
	return &Service{
		store:  storage,
		config: cfg,
	}
}

func (s *Service) generateHash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + s.config.PasswordSalt))
	return hex.EncodeToString(hasher.Sum(nil))
}

type claims struct {
	PersonId *int64 `json:"person_id,omitempty"`
	jwt.RegisteredClaims
}

func (s *Service) generateToken(claims claims, secret string) (*string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *Service) CreatePerson(username, password string) (*entity.Person, error) {

	if len(password) < 10 {
		return nil, ErrPasswordTooShort{}
	}

	person, err := s.store.CreatePerson(username, s.generateHash(password))

	if err != nil {
		var errAlreadyExists storage.ErrAlreadyExists
		if errors.As(err, &errAlreadyExists) {
			return nil, ErrUsernameAlreadyTaken{err}
		}

		return nil, ErrInternal{err}
	}

	return person, err
}

func (s *Service) GenerateTokens(personId int64) (*string, *string, error) {
	data := claims{}
	data.PersonId = &personId
	data.IssuedAt = jwt.NewNumericDate(time.Now())
	data.ExpiresAt = jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour))

	refreshToken, err := s.generateToken(data, s.config.JwtRefreshSecret)
	if err != nil {
		return nil, nil, ErrInternal{err}
	}

	data.ExpiresAt = jwt.NewNumericDate(time.Now().Add(15 * time.Minute))

	accessToken, err := s.generateToken(data, s.config.JwtAccessSecret)
	if err != nil {
		return nil, nil, ErrInternal{err}
	}

	return accessToken, refreshToken, nil
}

func (s *Service) Refresh(refreshToken string) (*string, *string, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidRefreshToken{errors.New("incorrect signing method")}
		}
		return []byte(s.config.JwtRefreshSecret), nil
	})
	if err != nil {
		return nil, nil, ErrInvalidRefreshToken{err}
	}

	if !token.Valid {
		return nil, nil, ErrInvalidRefreshToken{errors.New("token is not valid")}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, ErrInvalidRefreshToken{err}
	}

	personId, ok := claims["person_id"].(int64)
	if !ok {
		return nil, nil, ErrInvalidRefreshToken{err}
	}

	newAccessToken, newRefreshToken, err := s.GenerateTokens(personId)
	if err != nil {
		return nil, nil, err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *Service) SignIn(username, password string) (*entity.Person, error) {
	person, err := s.store.GetPersonByUsername(username)
	if err != nil {
		var ErrNotFound storage.ErrNotFound
		if errors.As(err, &ErrNotFound) {
			return nil, ErrInvalidCredentials{err}
		}

		return nil, ErrInternal{err}
	}

	if *person.Password != s.generateHash(password) {
		return nil, ErrInvalidCredentials{errors.New("password is incorrect")}
	}

	person.Password = nil

	return person, err
}
