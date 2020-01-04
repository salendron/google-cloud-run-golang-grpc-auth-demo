package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

// server is used to implement UnimplementedTestServerServer.
type server struct {
	UnimplementedTestServerServer
}

// Multiply implements Multiply
func (s *server) Multiply(ctx context.Context, in *MultiplyRequest) (*MultiplyReply, error) {
	return &MultiplyReply{Result: in.GetX() * in.GetY()}, nil
}

func main() {
	// PORT is being set by the Cloud Run environment
	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("Starting on Port: %v\n", port)

	s := grpc.NewServer()
	RegisterTestServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
