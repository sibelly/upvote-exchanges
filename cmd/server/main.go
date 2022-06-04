package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
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

	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	exchangeService := exchanges.NewExchangeServiceServer()
	exchangeService.LoadFeatures(*jsonDBFile)

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

		pb.RegisterExchangesServiceServer(baseServer, exchangeService)

		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)

}
