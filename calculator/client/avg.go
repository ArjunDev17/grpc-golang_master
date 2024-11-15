package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
)

func doAvg(c pb.SumServicesClient) {
	log.Println("do Avg was invoked")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while opening the stream :%v\n", err)
	}
	numbers := []int32{1, 2, 3, 4, 5}
	for _, number := range numbers {
		log.Printf("Sending Number :%d\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while reciving response :%v\n", err)
	}
	log.Printf("Result is :%v", res.Result)
}
