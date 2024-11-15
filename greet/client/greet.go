package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet Requst called")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Kabbu",
	})
	if err != nil {
		log.Fatalf("could' not greet :%v\n", err)
	}
	log.Printf("Greeting %s\n", res.Result)
}
