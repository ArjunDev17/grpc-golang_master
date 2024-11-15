package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
)

func dOPrime(c pb.SumServicesClient) {
	log.Println("do Prime was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}
	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes:%v\n ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream :%v\n", err)
		}

		log.Printf("primes :%d\n", res.Result)
	}

}
