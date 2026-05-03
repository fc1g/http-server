package server

import (
	"net"

	"github.com/fc1g/http-server/internal/config"
)

func HandleConnection(conn net.Conn, cfg config.Connection) {
	defer conn.Close()

	// buffer := make([]byte, cfg.BufferSize)

	for {
		// set read timeout
		//
		// read from connection
		//
		// parse request
		//
		// validate request
		//
		// set write timeout
		//
		// write into connection
	}
}
