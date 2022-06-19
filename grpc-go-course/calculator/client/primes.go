package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error was calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}
		log.Printf("Primes: %v\n", msg.Result)
	}
}
