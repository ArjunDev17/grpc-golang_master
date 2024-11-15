package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.Number + in.Number2,
	}, nil
}
