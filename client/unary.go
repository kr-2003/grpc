package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kr-2003/grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not be greeted: %v", err)
	}
	log.Printf("%s", res.Message)
}
