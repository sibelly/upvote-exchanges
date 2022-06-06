package exchanges

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/sibelly/upvote-exchanges/configs"
	"github.com/sibelly/upvote-exchanges/pb"
	"google.golang.org/grpc"
)

var mongoTestUri string = "mongodb+srv://sibelly:4020514ngc6302@cluster0.hysca.mongodb.net/?retryWrites=true&w=majority"

var (
	serviceServer *ExchangeServiceServer
)

func init() {
	err := os.Setenv("MONGOURI", mongoTestUri)

	// Connect to MongoDB
	_, err = configs.GetMongoClient()
	if err != nil {
		log.Fatalf("Could not connect to Mongo")
	}

	// Create gRPC server
	server := grpc.NewServer()
	serviceServer = NewExchangeServiceServer()
	pb.RegisterExchangesServiceServer(server, serviceServer)
}

// TestUpvote tests upvoting a Exchange record.
func TestUpvote(t *testing.T) {
	// Read previous value
	previous, err := serviceServer.ReadExchange(context.Background(), &pb.ReadReq{ExchangeId: "OKCOIN_CNY"})

	// Upvote exchange
	res, err := serviceServer.Upvote(context.Background(), &pb.VoteRequest{ExchangeId: "OKCOIN_CNY"})

	fmt.Println("res => ", res)

	if err != nil || res.ExchangeId != "OKCOIN_CNY" {
		t.Error("Error upvoting exchange")
	}

	// Read new value
	new, err := serviceServer.ReadExchange(context.Background(), &pb.ReadReq{ExchangeId: "OKCOIN_CNY"})

	if err != nil || previous.GetExchange().Upvotes > new.GetExchange().Upvotes {
		t.Error("Failed to upvote exchange", err)
	}
}

// // TestList tests listing the stream of exchanges
// func TestList(t *testing.T) {
// 	_, err := serviceServer.ListExchanges(context.Background(), &pb.Empty{})
// 	if err == nil {
// 		t.Error("Read an invalid exchange")
// 	}

// 	res, err := serviceServer.ReadExchange(context.Background(), &pb.ReadReq{ExchangeId: "OKCOIN_CNY"})
// 	if err != nil || res.GetExchange().ExchangeId != "OKCOIN_CNY" {
// 		t.Error("Could not read exchange")
// 	}
// }

// TestRead tests reading an unexisting exchange and an existing one
func TestRead(t *testing.T) {
	_, err := serviceServer.ReadExchange(context.Background(), &pb.ReadReq{ExchangeId: "invalid"})
	if err == nil {
		t.Error("Read an invalid exchange")
	}

	res, err := serviceServer.ReadExchange(context.Background(), &pb.ReadReq{ExchangeId: "OKCOIN_CNY"})
	if err != nil || res.GetExchange().ExchangeId != "OKCOIN_CNY" {
		t.Error("Could not read exchange")
	}
}
