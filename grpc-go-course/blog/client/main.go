package main

import (
	"fmt"
	"log"

	pb "github.com/Jaidip1994/GoLangHandsOn/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	// readBlog(c, "12")
	updateBlog(c, id)
	fmt.Println("Content after updating")
	readBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
