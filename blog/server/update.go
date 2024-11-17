package main

import (
	"context"

	pb "github.com/ArjunDev17/grpc-golang_master/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Can't parse ID",
		)
	}
	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
	}
	res, err := collection.UpdateOne(
		ctx, bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"could't update",
		)
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"can't find Blog with ID",
		)
	}
	return &emptypb.Empty{}, nil
}
