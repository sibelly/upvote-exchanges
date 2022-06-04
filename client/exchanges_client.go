package client

import (
	"context"
	"log"
	"time"

	"github.com/sibelly/upvote-exchanges/pb"
	"google.golang.org/grpc"
)

type ExchangeClient struct {
	service pb.ExchangesServiceClient
}

func NewExchangeClient(cc *grpc.ClientConn) *ExchangeClient {
	service := pb.NewExchangesServiceClient(cc)
	return &ExchangeClient{service}
}

func (exchangeClient *ExchangeClient) ListExchanges(req *pb.ExchangesServiceClient) {
	log.Printf("Received: %v", req)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := exchangeClient.service.ListExchange(ctx, &pb.Empty{})

	if err != nil {
		log.Fatalf("could not list exchanges: %v", err)
	}
	log.Printf("Exchanges: %s", res)

}
