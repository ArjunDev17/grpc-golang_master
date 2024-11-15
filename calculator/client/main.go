package main

import (
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:5059"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect :%v\n", err)
	}
	defer conn.Close()

	// c := pb.NewGreetServiceClient(conn)
	c := pb.NewSumServicesClient(conn)

	// doSum(c)
	// dOPrime(c)

	doAvg(c)
}
