package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
)

func doSum(c pb.SumServicesClient) {
	log.Println("doSum Requst called")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Number:  12,
		Number2: 3,
	})
	if err != nil {
		log.Fatalf("could' not greet :%v\n", err)
	}
	log.Printf("Sum is  %d\n", res.Result)
}
