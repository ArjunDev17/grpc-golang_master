package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

var addres string = "localhost:5050"

// arjun.singh@Update grpc-golang_master % evans --host localhost --port 5050 --reflection repl
func main() {
	//DataBase Connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("blogDB").Collection("blogCollection")
	//connection done
	//Server Connnectivity
	lis, err := net.Listen("tcp", addres)
	if err != nil {
		log.Fatalf("failded listen on :%v\n", err)
	}
	log.Printf("Listing on :%s\n", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n ", err)
	}
}
