package main

import (
	"context"
	"fmt"
	pb "github.com/pangpy/rpcfw/calculator.pb.go"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	result := in.A + in.B
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
