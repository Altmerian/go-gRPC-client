package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Altmerian/go-gRPC-client/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client instance.
	client := pb.NewPingPongServiceClient(conn)

	// Call the ping-pong method.
	req := &pb.PingRequest{
		Message:   "Ping",
		Timestamp: time.Now().UnixMilli(),
	}
	resp, err := client.PingPong(context.Background(), req)
	if err != nil {
		log.Fatalf("could not ping-pong: %v", err)
	}

	// Print the response.
	log.Printf("ping-pong response: %v", resp)
}
