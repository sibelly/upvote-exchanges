package client

import (
	"context"
	"log"
	"time"

	"github.com/sibelly/upvote-exchanges/pb"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type GreeterClient struct {
	service pb.GreeterClient
}

// NewGreeterClient returns a new helloworld client
func NewGreeterClient(cc *grpc.ClientConn) *GreeterClient {
	service := pb.NewGreeterClient(cc)
	return &GreeterClient{service}
}

// SayHello implements helloworld.GreeterServer
func (greeterClient *GreeterClient) SayHello(req *pb.HelloRequest) {
	log.Printf("Received: %v", req.GetName())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := greeterClient.service.SayHello(ctx, req)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.GetMessage())

}
