package httpController

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type personContextType string

const personContextKey personContextType = "person_ctx"

func jwtRead(jwtSecret string, log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Error("httpController:jwtRead", zap.Error(errors.New("missing Authorization header")))
				responseError(w, "Authorization required", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				log.Error("httpController:jwtRead", zap.Error(errors.New("invalid Authorization header")))
				responseError(w, "invalid Authorization header", http.StatusUnauthorized)
				return
			}
			tokenStr := parts[1]

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
				if t.Method != jwt.SigningMethodHS256 {
					return nil, errors.New("incorrect signing method")
				}
				return []byte(jwtSecret), nil
			})
			if err != nil {
				log.Error("httpController:jwtRead", zap.Error(err))
				responseError(w, "invalid Authorization header", http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				log.Error("httpController:jwtRead", zap.Error(errors.New("token is not valid")))
				responseError(w, "invalid Authorization header", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Error("httpController:jwtRead", zap.Error(errors.New("invalid refresh token structure")))
				responseError(w, "invalid Authorization header", http.StatusUnauthorized)
				return
			}

			personIdValue, ok := claims["person_id"]
			if !ok {
				log.Error("httpController:jwtRead", zap.Error(errors.New("token does not contain person_id")))
				responseError(w, "invalid Authorization header", http.StatusUnauthorized)
				return
			}

			personId, ok := personIdValue.(float64)
			if !ok {
				log.Error("httpController:jwtRead", zap.Error(errors.New("invalid person_id type")))
				responseError(w, "invalid Authorization header", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), personContextKey, int64(personId))
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func requestPerson(r *http.Request) *int64 {
	uc, ok := r.Context().Value(personContextKey).(int64)

	if !ok {
		return nil
	}

	return &uc
}
