package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	"github.com/sibelly/upvote-exchanges/configs"
	"github.com/sibelly/upvote-exchanges/endpoints"
	"github.com/sibelly/upvote-exchanges/pb"
	"github.com/sibelly/upvote-exchanges/service"
	transport "github.com/sibelly/upvote-exchanges/transports"
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
	fmt.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	flag.Parse()
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()
	// pb.RegisterGreeterServer(grpcServer, &server{})

	// // Reflection
	// reflection.Register(grpcServer)

	// log.Printf("server listening at %v", lis.Addr())
	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	//run database
	fmt.Printf("Running Databaseee!!!!!!!")
	client, err := configs.GetMongoClient()
	if err != nil {
		fmt.Printf("Errorrrr => %d", err)
	}
	//Create a handle to the respective collection in the database.
	collection := configs.GetCollection(client, "users")
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), collection)
	if err != nil {
		fmt.Printf("Errorrrr 22 => %d", err)
	}
	/////////

	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	addservice := service.NewService(logger)
	addendpoint := endpoints.MakeEndpoints(addservice)
	grpcServer := transport.NewGRPCServer(addendpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		// Reflection
		reflection.Register(baseServer)

		pb.RegisterMathServiceServer(baseServer, grpcServer)
		pb.RegisterGreeterServer(baseServer, &server{})

		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)

}
