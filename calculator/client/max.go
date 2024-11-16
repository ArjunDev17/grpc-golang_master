package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
)

func doMax(c pb.SumServicesClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error while opening stream :%v\n", err)
	}

	waitCh := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {

			log.Printf("sending number : %d\n", numbers)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
			log.Printf("Recived a new maximum :%d\n", res.Result)
		}
		close(waitCh)
	}()
	<-waitCh
}
