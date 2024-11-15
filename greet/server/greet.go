package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greeet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "hello " + in.FirstName,
	}, nil
}
