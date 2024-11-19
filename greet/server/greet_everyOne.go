package main

import (
	"io"
	"log"

	pb "github.com/ArjunDev17/grpc-golang_master/greet/proto"
)

// GreetEveryone should use the correct signature for bidirectional streaming
func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone Function was invoked")

	// Handle the stream here
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Client is done sending messages
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving from client: %v\n", err)
			return err
		}

		log.Printf("Received: %v\n", req.FirstName)

		// Send a response back to the client
		res := &pb.GreetResponse{
			Result: "Hello " + req.FirstName,
		}

		// Send the response to the client
		if err := stream.Send(res); err != nil {
			log.Fatalf("Error while sending response to client: %v\n", err)
			return err
		}
	}
}
