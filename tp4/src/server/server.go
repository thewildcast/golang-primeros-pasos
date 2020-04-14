package server

import (
	"net/http"
	v1 "server/api/v1"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouteHandler() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Mount("/api/v1/", v1.GetApiHandler())

	return router
}
