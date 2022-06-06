package main

import (
	"flag"
	"log"

	"github.com/sibelly/upvote-exchanges/client"
	"github.com/sibelly/upvote-exchanges/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client.Upvote(&client.Conn{Cc: conn}, nil)

	client.ListExchanges(&client.Conn{Cc: conn}, &pb.Empty{})

}
