package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("sqrt was invoked with: %v\n", in)
	number := float64(in.Number)
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %f", number),
		)
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(number),
	}, nil
}
