package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/fc1g/http-server/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s, err := server.New()
	if err != nil {
		log.Fatal("Failed to initialize server: ", err)
	}

	go func() {
		if err := s.Run(); err != nil {
			log.Fatal("Failed to run server: ", err)
		}
	}()

	<-ctx.Done()

	if err := s.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
}
