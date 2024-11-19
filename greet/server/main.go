package main

import (
	"log"
	"net"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	pb.GreetServiceServer
}

var addres string = "localhost:5050"

func main() {
	lis, err := net.Listen("tcp", addres)

	opts := []grpc.ServerOption{}
	if err != nil {
		log.Fatalf("failded listen on :%v\n", err)
	}
	log.Printf("Listing on :%s\n", lis.Addr())
	tls := true
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading Certificate :%v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n ", err)
	}
}
