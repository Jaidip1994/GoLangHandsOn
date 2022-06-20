package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Jai"},
		{FirstName: "Sash"},
		{FirstName: "SaJa"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling long greet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending Req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response %v\n", err)
	}

	log.Printf("LongGreet: %v\n", res.Result)

}
