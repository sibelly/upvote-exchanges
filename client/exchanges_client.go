package client

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/sibelly/upvote-exchanges/pb"
	"google.golang.org/grpc"
)

type Conn struct {
	Cc *grpc.ClientConn
}

type ExchangeClient struct {
	pb.ExchangesServiceClient
}

func NewExchangeClient(cc *grpc.ClientConn) *ExchangeClient {
	service := pb.NewExchangesServiceClient(cc)
	return &ExchangeClient{service}
}

func Upvote(conn *Conn, req *pb.VoteRequest) (res *pb.VoteResponse) {
	log.Printf("Received Upvote: %v", req)

	service := NewExchangeClient(conn.Cc)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := service.Upvote(ctx, &pb.VoteRequest{ExchangeId: "HUOBI"})

	if err != nil {
		log.Fatalf("could not list exchanges: %v", err)
	}
	log.Printf("Exchanges: %s", res)

	return res

}

func ListExchanges(conn *Conn, req *pb.Empty) (res *pb.ExchangesService_ListExchangesClient) {
	log.Printf("Received ListExchanges: %v", req)

	service := NewExchangeClient(conn.Cc)

	ctx := context.TODO()

	in := &pb.Empty{Id: 1}

	stream, err := service.ListExchanges(ctx, in)

	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")

	return res

}
