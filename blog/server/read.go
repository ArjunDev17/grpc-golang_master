package main

import (
	"context"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("Read Blog invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Can't parse ID",
		)
	}
	data := &BlogItem{}

	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot find blog with the ID provider",
		)
	}
	return documentToBlog(data), nil
}
