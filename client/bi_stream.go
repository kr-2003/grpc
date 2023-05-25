package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kr-2003/grpc/proto"
)

func callHelloBidirecttionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("bidirectional streaming has started")
	stream, err := client.SayHelloBidirecttionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending the request %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("bidirectional streaming finished")
}
