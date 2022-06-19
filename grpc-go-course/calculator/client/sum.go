package main

import (
	"context"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	})

	if err != nil {
		log.Printf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Response)
}
