package main

import (
	"log"
	"time"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
	"google.golang.org/grpc"
)

// Implement the GreetManyTimes method (Server-streaming RPC)
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	// Send a few responses in a loop, simulating a long-running stream
	for i := 0; i < 5; i++ {
		// Create a greeting message for the current iteration
		result := "Hello " + in.FirstName + " #" + string(i+1)

		// Send the response
		if err := stream.SendMsg(&pb.GreetResponse{Result: result}); err != nil {
			log.Fatalf("Error sending message: %v\n", err)
			return err
		}

		// Simulate some delay (e.g., waiting for some server-side process)
		time.Sleep(1 * time.Second)
	}

	// Once the stream is complete, return nil (which will close the stream on the server side)
	return nil
}
