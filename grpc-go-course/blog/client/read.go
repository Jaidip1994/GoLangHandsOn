package main

import (
	"context"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")
	req := &pb.BlogId{
		Id: id,
	}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Unexpected error occurred: %v\n", err)
	}
	log.Printf("Blog was read: %v\n", res)
	return res
}
