package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/fc1g/http-server/internal/config"
	"github.com/fc1g/http-server/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server, err := server.New(cfg.Server)
	if err != nil {
		log.Fatal("Failed to initialize server: ", err)
	}

	go func() {
		if err := server.Run(); err != nil {
			log.Fatal("Failed to run server: ", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
}
