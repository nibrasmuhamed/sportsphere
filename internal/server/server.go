package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nibrasmuhamed/sportsphere/config"
	"github.com/nibrasmuhamed/sportsphere/internal/middleware"
	"github.com/nibrasmuhamed/sportsphere/internal/routes"
	"github.com/nibrasmuhamed/sportsphere/pkg/logger"
	"go.uber.org/zap"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()
	middleware.RegisterMiddleware(r)

	routes.RegisterRoutes(r)

	server := &http.Server{
		Addr:         ":" + config.Get().Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		logger.Info("Starting server on port " + config.Get().Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error starting server", zap.Error(err))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	select {
	case <-stop:
		logger.Info("Shutdown signal received, starting graceful shutdown...")
	case <-ctx.Done():
		logger.Info("Shutdown initiated by external context")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("Error during shutdown", zap.Error(err))
	}

	logger.Info("Server stopped gracefully.")
}
