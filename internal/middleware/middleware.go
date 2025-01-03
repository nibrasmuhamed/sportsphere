package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// RegisterMiddleware is used to apply common middleware functions to the router.
func RegisterMiddleware(r *chi.Mux) {

	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

}
