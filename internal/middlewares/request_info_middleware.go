package middlewares

import (
	"context"
	"net/http"

	context_key "gungun974.com/viteintegration/internal/context"
)

type AppRequestInfo struct {
	ActiveUrl string
}

func RequestInfoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(
			ctx,
			context_key.REQUEST_INFO_KEY,
			AppRequestInfo{ActiveUrl: r.RequestURI},
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
