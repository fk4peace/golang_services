package httpController

import (
	"net/http"

	"go.uber.org/zap"
)

func logger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			next.ServeHTTP(w, r)

			log.Info("http request", zap.String("method", r.Method), zap.String("path", r.URL.Path))
		})
	}
}
