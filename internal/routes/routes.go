package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mohdjishin/sportsphere/internal/handlers"
	"github.com/mohdjishin/sportsphere/internal/repository"
	"github.com/mohdjishin/sportsphere/pkg/service"
)

func RegisterRoutes(r *chi.Mux) {
	registerAPIRoutes(r)
}

func registerAPIRoutes(r chi.Router) {
	// an abstraction to handle health check
	healthCheck := handlers.NewHealthCheck()
	operatorRepository := repository.NewOperatorRepository()
	operatorService := service.NewOperatorService(operatorRepository)
	operatorController := handlers.NewOperatorHandler(operatorService)
	//operator routes
	r.Post("/api/v1/operator", operatorController.CreateOperator)

	r.Get("/health", healthCheck.HealthCheck)

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
