package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sibelly/upvote-exchanges/configs"
	"github.com/sibelly/upvote-exchanges/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//run database
	log.Printf("Running Databaseee!!!!!!!")
	client, err := configs.GetMongoClient()
	if err != nil {
		log.Fatalln("Errorrrr => ", err)
	}
	//Create a handle to the respective collection in the database.
	collection := configs.GetCollection(client, "users")
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), collection)
	if err != nil {
		log.Fatalln("Errorrrr 22 => ", err)
	}
	//Return success without any error.
	/////////

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	// Reflection
	reflection.Register(grpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
