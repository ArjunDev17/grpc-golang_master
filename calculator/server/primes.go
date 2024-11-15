package main

import (
	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(ctx *pb.PrimeRequest, in grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	number := ctx.Number
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			in.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}

	}
	return nil
}
