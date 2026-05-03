package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/fc1g/http-server/internal/config"
	"github.com/fc1g/http-server/internal/server"
	"github.com/fc1g/http-server/internal/validation"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const configFilePath = "config.yaml"

func init() {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading %s", configFilePath)
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	v := validation.New()

	cfg, err := config.Load(v)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	_, err = config.LoadSecrets(v)
	if err != nil {
		log.Fatalf("Failed to load secrets: %v", err)
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
