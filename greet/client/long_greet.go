package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("do long Greet was invoked")
	regs := []*pb.GreetRequest{
		{FirstName: "Ram"},
		{FirstName: "Shyam"},
		{FirstName: "Raghu"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while calling LongGreet %v\n", err)
	}
	for _, req := range regs {
		log.Printf(" sending req :%v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reciving response from LongGreet :%v\n", err)
	}
	log.Printf("LongGreet :%s\n", res.Result)
}
