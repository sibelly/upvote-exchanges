package exchanges

import (
	"context"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/sibelly/upvote-exchanges/pb"
)

type gRPCServer struct {
	upvote gt.Handler
	// listExchanges gt.Handler
	pb.UnimplementedExchangesServiceServer
}

// NewGRPCServer initializes a new gRPC server
func NewExchangeGRPCServer(endpoints Endpoints, logger log.Logger) pb.ExchangesServiceServer {
	return &gRPCServer{
		upvote: gt.NewServer(
			endpoints.Upvote,
			decodeVoteRequest,
			encodeVoteResponse,
		),
		// listExchanges: gt.NewServer(
		// 	endpoints.ListExchange,
		// 	decodeEmptyRequest,
		// 	encodeListResponse,
		// ),
	}
}

func (s *gRPCServer) Upvote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	_, resp, err := s.upvote.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.VoteResponse), nil
}

// func (s *gRPCServer) ListExchange(ctx context.Context, req *pb.Empty) ([]*pb.ListResponse, error) {
// 	_, resp, err := s.listExchanges.ServeGRPC(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.([]*pb.ListResponse), nil
// }

func decodeVoteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.VoteRequest)
	return VoteReq{Exchange_id: req.ExchangeId}, nil
}

func encodeVoteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(VoteResp)
	return &pb.VoteResponse{Name: resp.Name, Website: resp.Website, Votes: resp.Votes}, nil
}

// func decodeEmptyRequest(_ context.Context, request interface{}) (interface{}, error) {
// 	req := request.(*pb.Empty)
// 	return VoteReq{req.String()}, nil
// }

// func encodeListResponse(_ context.Context, response interface{}) (interface{}, error) {
// 	resp := response.(ExchangeType)
// 	return resp, nil
// }
