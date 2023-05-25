package main

import (
	"io"
	"log"

	pb "github.com/kr-2003/grpc/proto"
)

func (s *helloServer) SayHelloBidirecttionalStreaming(stream pb.GreetService_SayHelloBidirecttionalStreamingServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
