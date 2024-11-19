package main

import (
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address string = "localhost:5050"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("error while loading CA certificat:%v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("Failed to connect :%v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetWithDeadline(c, 5*time.Second)
	doGreetEveryone(c)

}
