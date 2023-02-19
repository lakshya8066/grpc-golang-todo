package main

import (
	"context"
	"io"
	"log"

	pb "github.com/lakshya8066/grpc-golang-server/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started!")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

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
	log.Printf("Streaming finished")
}
