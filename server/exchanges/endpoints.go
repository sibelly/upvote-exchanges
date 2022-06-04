package exchanges

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	Upvote       endpoint.Endpoint
	ListExchange endpoint.Endpoint
}

type VoteReq struct {
	Exchange_id string
}

type VoteResp struct {
	Name    string
	Website string
	Votes   int32
}

type EmptyReq struct{}

type ListResp struct {
	Exchanges []ExchangeType
	// string id = 1;
	// string exchange_id = 2;
	// string website = 3;
	// string name = 4;
	// int32 upvotes = 5;
	// int32 downvotes = 6;
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s ExchangeService) Endpoints {
	return Endpoints{
		Upvote: makeUpvoteEndpoint(s),
		// ListExchange: makeListExchangeEndpoint(s),
	}
}

func makeUpvoteEndpoint(s ExchangeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(VoteReq)
		result, _ := s.Upvote(ctx, req.Exchange_id)
		return VoteResp{Name: result.Name, Website: result.Website, Votes: result.Upvotes}, nil
	}
}

// func makeListExchangeEndpoint(s ExchangeService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		result, _ := s.ListExchange(ctx)
// 		return ListResp{result}, nil
// 	}
// }
