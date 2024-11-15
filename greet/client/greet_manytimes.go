package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreet Requst called")
	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Kabbu",
	})
	if err != nil {
		log.Fatalf("could' not greet many times :%v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream :%v\n ", err)
		}
		log.Printf("Greeting many times :%s\n", msg.Result)
	}

}
