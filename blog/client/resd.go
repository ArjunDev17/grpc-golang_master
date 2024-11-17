package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("read blog invoked for Id :%s\n", id)

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("error hapend while reading :%v", err)
	}
	log.Printf("Blog was read :%v\n", res)
	return res
}
