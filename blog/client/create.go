package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--createBlog was invoked--")
	blog := &pb.Blog{
		AuthorId: "108",
		Title:    "Jap",
		Content:  "Content of the First blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected Error :%v\n", err)
	}
	log.Printf("Blog has created :%s\n", res.Id)
	return res.Id
}
