package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func ConditionalLogger(paths []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, path := range paths {
				if r.URL.Path == path {
					next.ServeHTTP(w, r)
					return
				}
			}
			middleware.Logger(next).ServeHTTP(w, r)
		})
	}
}
