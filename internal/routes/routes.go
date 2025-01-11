package routes

// Swagger URL: http://localhost:8000/swagger/index.html

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/nibrasmuhamed/sportsphere/db"
	_ "github.com/nibrasmuhamed/sportsphere/docs" // Import for Swagger documentation
	"github.com/nibrasmuhamed/sportsphere/internal/handlers"
	"github.com/nibrasmuhamed/sportsphere/internal/repository"
	"github.com/nibrasmuhamed/sportsphere/pkg/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Sportsphere API
// @version 1.0
// @description This is a sample API for the Sportsphere platform.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1

func RegisterRoutes(r *chi.Mux) {
	registerAPIRoutes(r)
}

func registerAPIRoutes(r chi.Router) {
	// an abstraction to handle health check
	healthCheck := handlers.NewHealthCheck()
	uow, _ := db.NewUnitOfWork(context.Background())
	operatorRepository := repository.NewOperatorRepository()
	operatorService := service.NewOperatorService(operatorRepository, uow)
	operatorController := handlers.NewOperatorHandler(operatorService)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	//operator routes
	r.Post("/api/v1/operator", operatorController.CreateOperator)

	r.Get("/health", healthCheck.HealthCheck)

}
