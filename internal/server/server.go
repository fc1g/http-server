package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
}

const network = "tcp"
const port = ":3000"

func New() (*Server, error) {
	l, err := net.Listen(network, port)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection on port %s, %w", port, err)
	}

	return &Server{
		listener: l,
	}, nil
}

func (s *Server) Run() error {
	log.Printf("Server is listening on %s\n", port)

	return nil
}

func (s *Server) Shutdown() error {
	if err := s.listener.Close(); err != nil {
		return fmt.Errorf("failed to close listener: %w", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
