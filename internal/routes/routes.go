package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mohdjishin/sportsphere/internal/handlers"
)

func RegisterRoutes(r *chi.Mux) {
	registerAPIRoutes(r)

}

func registerAPIRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		registerInfoRoutes(r)
		registerUserRoutes(r)
	})
}

func registerInfoRoutes(r chi.Router) {
	r.Get("/info", handlers.Info)
}

func registerUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsers)
		r.Post("/", handlers.CreateUser)
	})
}
