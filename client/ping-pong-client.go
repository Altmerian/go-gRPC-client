package client

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"

    pb "github.com/Altmerian/go-gRPC-client/proto"
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
        Message:   "hello",
        Timestamp: time.Now().UnixNano(),
    }
    resp, err := client.PingPong(context.Background(), req)
    if err != nil {
        log.Fatalf("could not ping-pong: %v", err)
    }

    // Print the response.
    log.Printf("ping-pong response: %v", resp)
}
