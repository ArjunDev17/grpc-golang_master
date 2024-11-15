package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {

	log.Println("Long Greet Function was invoked")

	res := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading Client stream %v\n", err)
		}
		log.Printf("Reciving :%v\n", req)
		res += fmt.Sprintf("Hello %s:\n", req.FirstName)

	}

}
