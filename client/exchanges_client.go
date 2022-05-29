package client

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/sibelly/upvote-exchanges/pb"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

// GreeterClient is a client to call helloworld service RPCs
type GreeterClient struct {
	service pb.GreeterClient
}

// NewGreeterClient returns a new helloworld client
func NewGreeterClient(cc *grpc.ClientConn) *GreeterClient {
	service := pb.NewGreeterClient(cc)
	return &GreeterClient{service}
}

// SayHello calls create helloworld RPC
func (greeterClient *GreeterClient) SayHello(req *pb.HelloRequest) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := greeterClient.service.SayHello(ctx, &pb.HelloRequest{Name: req.Name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = greeterClient.service.SayHelloAgain(ctx, &pb.HelloRequest{Name: req.Name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
