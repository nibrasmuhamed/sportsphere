package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mohdjishin/sportsphere/config"
	"github.com/mohdjishin/sportsphere/internal/server"
	"github.com/mohdjishin/sportsphere/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Run(ctx)

	go server.Run(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Info("Received shutdown signal. Initiating graceful shutdown.")

	cancel()
}
