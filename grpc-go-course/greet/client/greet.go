package main

import (
	"context"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Jaidip",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %v\n", res.Result)
}
