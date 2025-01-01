package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mohdjishin/sportsphere/config"
	"github.com/mohdjishin/sportsphere/logger"
	"go.uber.org/zap"
)

// StartServer initializes and starts the server
func StartServer() {
	// initialising configuration
	config.Init()
	// initialising logger
	logger.Init()

	// router setup
	r := chi.NewRouter()

	// Global middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Custom middleware for handling CORS (optional)
	r.Use(corsMiddleware)

	// Routes
	r.Mount("/api/v1", apiRouter())

	// Configurable server parameters
	server := &http.Server{
		Addr:         ":" + config.Config.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown handling
	go func() {
		logger.Logger.Info("Starting server on port " + config.Config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("Error starting server: %v", zap.Error(err))
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logger.Logger.Info("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error("Error during shutdown: %v", zap.Error(err))
	}

	logger.Logger.Info("Server stopped")
}

// Example modular API router
func apiRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", listUsers)
		r.Post("/", createUser)
	})

	return r
}

// Example handlers
func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a user"))
}

// Example custom middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	StartServer()
}
