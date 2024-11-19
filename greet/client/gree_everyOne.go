package main

import (
	"context"
	"io"
	"log"

	"time"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("DoGreet everyone invoked")
	stream, err := c.GreetManyTimes(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream :%v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Ram Ji"},
		{FirstName: "Shyam Ji"},
		{FirstName: " Ji"},
		{FirstName: "hanuman Ji"},
	}
	waitCh := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending Request :%v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error while reciving :%v\n", err)
				break
			}
			log.Printf("Recived :%v\n", res.Result)
		}
		close(waitCh)
	}()
	<-waitCh
}
