package server

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/fc1g/http-server/internal/config"
)

type Server struct {
	config   config.ServerConfig
	listener net.Listener
}

func New(cfg config.ServerConfig) (*Server, error) {
	listener, err := net.Listen(cfg.Network, fmt.Sprintf(":%d", cfg.Addr))
	if err != nil {
		return nil, fmt.Errorf("failed to open connection on port %d: %w", cfg.Addr, err)
	}

	return &Server{
		config:   cfg,
		listener: listener,
	}, nil
}

func (s *Server) Run() error {
	log.Printf("Server is listening on :%d\n", s.config.Addr)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}

			log.Printf("Failed to establish connection %v", err)
			continue
		}

		go func(net.Conn) {}(conn)
	}

	return nil
}

func (s *Server) Shutdown() error {
	if err := s.listener.Close(); err != nil {
		return fmt.Errorf("failed to close listener: %w", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
