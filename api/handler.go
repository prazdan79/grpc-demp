package api

import (
		"log"
		"context"
	   )

// Server represents the gRPC server
type Server struct {
}


// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
    log.Printf("Receive message from Exium Server: %s", in.Greetings)
    return &PingMessage{Greetings: "Hello Exium"}, nil
}
