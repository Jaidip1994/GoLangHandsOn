package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 24, 23}
	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving response: %v\n", err)
	}
	log.Printf("Avg: %f\n", res.Result)
}
