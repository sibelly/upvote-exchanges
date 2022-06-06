package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sibelly/upvote-exchanges/configs"
	"github.com/sibelly/upvote-exchanges/pb"
	"github.com/sibelly/upvote-exchanges/server/exchanges"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port       = flag.Int("port", 50051, "The server port")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of exchanges")
)

func main() {
	flag.Parse()

	// Connect to MongoDB
	mongoClient, err := configs.GetMongoClient()
	if err != nil {
		log.Fatalf("Could not initialize Mongo client: %v", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	// Create listen socket
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}

	// Create gRPC server
	server := grpc.NewServer()

	// Reflection
	reflection.Register(server)

	exchangeService := exchanges.NewExchangeServiceServer()
	exchangeService.LoadFeatures(*jsonDBFile)

	pb.RegisterExchangesServiceServer(server, exchangeService)
	log.Printf("Listening on %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}
