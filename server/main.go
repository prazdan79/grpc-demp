package main

import (
		"fmt"
		"log"
		"net"
	    "github.com/prazdan79/grpc-demp/api"
	    "google.golang.org/grpc"
)



func main () {
	lis, err :=  net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listern : %v", err)
	}

	s  := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
