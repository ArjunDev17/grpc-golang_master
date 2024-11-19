package main

import (
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:5050"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect :%v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	UpdateBlog(c, id)
	listBlog(c)
	deletBlog(c, id)
}
