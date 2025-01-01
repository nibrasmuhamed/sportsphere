package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/mohdjishin/sportsphere/internal/handlers"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/info", handlers.Info)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListtUsers)
		r.Post("/", handlers.CreateUser)
	})

	return r
}
