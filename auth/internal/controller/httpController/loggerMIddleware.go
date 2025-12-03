package httpController

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type errorType string

const errorKey errorType = "error_ctx"

func getError(ctx context.Context) error {
	v := ctx.Value(errorKey)
	if err, ok := v.(error); ok {
		return err
	}
	return nil
}

func writeError(r *http.Request, err error) {
	ctx := context.WithValue(r.Context(), errorKey, err)
	*r = *r.WithContext(ctx)
}

func logger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			next.ServeHTTP(w, r)
			err := getError(r.Context())

			log.Info("http request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

			if err == nil {
				return
			}
			log.Error(err.Error(), zap.Error(err))
		})
	}
}
