package main

import (
	"context"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
