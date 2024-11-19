package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("Do Greet Deadline invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req := &pb.GreetRequest{
		FirstName: "RamJi",
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		//basically when error from grpc
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exeeded ")
				return
			} else {
				log.Fatalf("unexpected grpc error :%v\n", err)
			}
		} else {
			log.Fatalf("A non grpc error :%v\n", err)
		}
	}
	log.Printf("Greet with dead line %s\n", res.Result)
}
