package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
)

func UpdateBlog(c pb.BlogServiceClient, id string) {
	log.Println("upfate blog was created")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "JaiRam",
		Title:    "Ramayna",
		Content:  "Maryada Purushotmm Raam",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happend with updating :%v\n", err)
		return
	}
	log.Println("Log was created")

}
