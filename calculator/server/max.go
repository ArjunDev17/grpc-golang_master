package main

import (
	"io"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	log.Println("Max server called")
	var maximum int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading clint stream :%v\n", err)
		}
		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})
			if err != nil {
				log.Fatalf("Error while sednig data  to client")
			}

		}
	}
}
