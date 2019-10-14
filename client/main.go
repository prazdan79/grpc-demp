package main

import (
		"log"
		"github.com/prazdan79/grpc-demp/api"
		"golang.org/x/net/context"
		"google.golang.org/grpc"
)

func main () {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greetings: "Hello Exium:"})
	if err != nil {
		log.Fatalf("Ping send failed %v", err)
	}
	log.Printf("Response from Exium server: %s", response.Greetings)
}


