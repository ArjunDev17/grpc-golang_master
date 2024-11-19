package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
)

func deletBlog(c pb.BlogServiceClient, id string) {
	log.Println("--delete Blog was invoked")
	_, err := c.DeleteingBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("error while deleteing :%v\n", err)
	}
	log.Println("Blog wwas created ")
}
