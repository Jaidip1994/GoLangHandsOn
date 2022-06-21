package main

import (
	"io"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/greet/proto"
)

func (s *Server) GreetEverone(stream pb.GreetService_GreetEveroneServer) error {
	log.Println("GreetEverone was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		err = stream.Send(&pb.GreetResponse{
			Result: "Hello " + req.FirstName + "!",
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
