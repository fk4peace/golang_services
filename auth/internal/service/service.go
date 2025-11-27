package service

import (
	"auth_service/internal/config"
	"auth_service/internal/entity"
	"auth_service/internal/storage"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type istorage interface {
	CreatePerson(username, password string) (*entity.Person, error)
	GetPersonByUsername(username string) (*entity.Person, error)
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

func (s *Service) passwordHash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + s.config.PasswordSalt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *Service) CreatePerson(username, password string) (*entity.Person, error) {

	if len(password) < 10 {
		return nil, ErrPasswordTooShort{}
	}

	person, err := s.store.CreatePerson(username, s.passwordHash(password))

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

	accessClaims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(personId, 10),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccess, err := accessToken.SignedString([]byte(s.config.JwtAccessSecret))
	if err != nil {
		return nil, nil, ErrInternal{err}
	}

	refreshClaims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(personId, 10),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefresh, err := refreshToken.SignedString([]byte(s.config.JwtRefreshSecret))
	if err != nil {
		return nil, nil, ErrInternal{err}
	}

	return &signedAccess, &signedRefresh, nil
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

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, nil, ErrInvalidRefreshToken{err}
	}

	userID, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return nil, nil, ErrInvalidRefreshToken{err}
	}

	newAccessToken, newRefreshToken, err := s.GenerateTokens(userID)
	if err != nil {
		return nil, nil, ErrInternal{err}
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

	if *person.Password != s.passwordHash(password) {
		return nil, ErrInvalidCredentials{errors.New("password is incorrect")}
	}

	person.Password = nil

	return person, err
}
