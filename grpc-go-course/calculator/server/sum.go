package main

import (
	"context"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v", in)
	return &pb.SumResponse{
		Response: in.FirstNumber + in.SecondNumber,
	}, nil
}
