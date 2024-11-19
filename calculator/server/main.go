package main

import (
	"log"
	"net"

	pb "github.com/ArjunDev17/grpc-golang_master/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedSumServicesServer
}

var addres string = "localhost:5050"

// arjun.singh@Update grpc-golang_master % evans --host localhost --port 5050 --reflection repl
func main() {
	lis, err := net.Listen("tcp", addres)
	if err != nil {
		log.Fatalf("failded listen on :%v\n", err)
	}
	log.Printf("Listing on :%s\n", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterSumServicesServer(s, &Server{})

	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n ", err)
	}
}
