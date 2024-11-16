package main

import (
	"log"
	"net"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSumServicesServer
}

var addres string = "localhost:5050"

func main() {
	lis, err := net.Listen("tcp", addres)
	if err != nil {
		log.Fatalf("failded listen on :%v\n", err)
	}
	log.Printf("Listing on :%s\n", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterSumServicesServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n ", err)
	}
}
