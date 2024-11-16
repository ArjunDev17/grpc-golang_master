package main

import (
	"io"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Avg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {
	log.Println("avg function called ")
	var sum int32 = 0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("error while reading Client Stream :%v\n ", err)
		}

		log.Printf("Reciving Number :%d\n", req.Number)
		sum += req.Number
		count++
	}

}
