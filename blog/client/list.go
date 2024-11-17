package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("List blogs invoked ")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("error whilie calling list block :%v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something wrong :%v\n", err)
		}
		log.Println(res)
	}
}
