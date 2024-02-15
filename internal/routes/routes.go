package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"gungun974.com/viteintegration/internal/middlewares"
)

func MainRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.StripSlashes)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middlewares.ConditionalLogger([]string{}))

	router.Use(middleware.Recoverer)

	router.Use(middlewares.ViteMiddleware)
	router.Use(middlewares.RequestInfoMiddleware)

	router.NotFound(HandleNotFoundPage)

	router.Mount("/", IndexRouter())

	return router
}
