package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nibrasmuhamed/sportsphere/config"
	_ "github.com/nibrasmuhamed/sportsphere/config"
	"github.com/nibrasmuhamed/sportsphere/db"
	"github.com/nibrasmuhamed/sportsphere/internal/server"
	"github.com/nibrasmuhamed/sportsphere/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config.Init()

	logger.Run(ctx)
	db.InitDatabase()

	go server.Run(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Info("Received shutdown signal. Initiating graceful shutdown.")

	cancel()
}
