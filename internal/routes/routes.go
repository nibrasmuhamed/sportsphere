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
// @description API for the Sportsphere platform.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1

func RegisterRoutes(r *chi.Mux) {
	uow, err := db.NewUnitOfWork(context.Background())
	if err != nil {
		panic("failed to initialize unit of work")
	}

	registerSwaggerRoutes(r)
	registerHealthRoutes(r)
	registerAPIV1Routes(r, uow)
}

func registerSwaggerRoutes(r chi.Router) {
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}

func registerHealthRoutes(r chi.Router) {
	r.Get("/health", handlers.NewHealthCheck().HealthCheck)
}

func registerAPIV1Routes(r chi.Router, uow db.UnitOfWork) {
	v1 := chi.NewRouter()
	r.Mount("/api/v1", v1)

	registerOperatorRoutes(v1, uow)
	registerUserRoutes(v1, uow)
}

func registerOperatorRoutes(r chi.Router, uow db.UnitOfWork) {
	repo := repository.NewOperatorRepository()
	service := service.NewOperatorService(repo, uow)
	controller := handlers.NewOperatorHandler(service)

	r.Post("/operator", controller.CreateOperator)
}

func registerUserRoutes(r chi.Router, uow db.UnitOfWork) {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo, uow)
	controller := handlers.NewUserController(service)

	r.Post("/user", controller.RegisterUser)
}
