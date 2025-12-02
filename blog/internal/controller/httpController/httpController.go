package httpController

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/fk4peace/golang_services/blog/internal/entity"
	"github.com/golang-jwt/jwt/v5"

	"github.com/go-chi/chi/v5"
)

type iService interface {
	GetPosts(limit, page int64) ([]entity.Post, *int64, error)
	GetPostById(postId int64) (*entity.Post, error)
	CreatePost(personId int64, content string) (*entity.Post, error)
}

type httpController struct {
	service   iService
	jwtSecret string
}

func New(service iService, jwtSecret string) http.Handler {
	controller := httpController{
		service,
		jwtSecret,
	}

	router := chi.NewRouter()

	router.Get("/posts", controller.getPosts)
	router.Get("/posts/{id}", controller.getPostById)

	router.Group(func(router chi.Router) {
		router.Use(JwtRead(controller.jwtSecret))
		router.Post("/posts", controller.createPost)

	})

	return router
}

type personContextType string

const personContextKey personContextType = "person_ctx"

func JwtRead(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				responseError(w, "missing Authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				responseError(w, "invalid Authorization header", http.StatusUnauthorized)
				return
			}
			tokenStr := parts[1]

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				responseError(w, "invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				responseError(w, "invalid token", http.StatusUnauthorized)
				return
			}

			var personId int64
			switch v := claims["person_id"].(type) {
			case float64:
				personId = int64(v)
			case int64:
				personId = v
			case int:
				personId = int64(v)
			case string:
				var parsed int64
				_, err := fmt.Sscan(v, &parsed)
				if err == nil {
					personId = parsed
				} else {
					responseError(w, "invalid token", http.StatusUnauthorized)
					return
				}
			default:
				responseError(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), personContextKey, personId)
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
