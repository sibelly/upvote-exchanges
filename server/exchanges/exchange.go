package exchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	otherLog "log"

	"github.com/go-kit/kit/log"
	"github.com/sibelly/upvote-exchanges/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type exchangeService struct {
	logger        log.Logger
	savedFeatures []*pb.Exchange // read-only after initialized
}

// ExchangeType represents the bson readable data from the protobuf
type ExchangeType struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Exchange_id string             `bson:"exchange_id"`
	Name        string             `bson:"name"`
	Website     string             `bson:"website"`
	Upvotes     int32              `bson:"upvotes"`
	Downvotes   int32              `bson:"downvotes"`
}

// Service interface
type ExchangeService interface {
	Upvote(ctx context.Context, exchange_id string) (ExchangeType, error)
	// Downvote(exchange_id string) (name string, website string, votes int32)

	// CreateExchange(website string, name string, upvotes int32, downvotes int32) ExchangeType
	// ReadExchange(exchange_id string) (website string, name string, upvotes int32, downvotes int32)
	ListExchange(ctx context.Context) ([]ExchangeType, error)
}

// NewService func initializes a service
func NewExchangeService(logger log.Logger) ExchangeService {
	return &exchangeService{
		logger: logger,
	}
}

func (s exchangeService) Upvote(ctx context.Context, exchange_id string) (ExchangeType, error) {
	fmt.Println("hueheuhedasdasdu")
	return ExchangeType{Name: "teste"}, nil
}

// func (s service) Downvote(ctx context.Context, numA, numB float32) (float32, error) {
// 	return numA - numB, nil
// }

// func (s service) CreateExchange(ctx context.Context, numA, numB float32) (float32, error) {
// 	return numA * numB, nil
// }

// func (s service) ReadExchange(ctx context.Context, numA, numB float32) (float32, error) {
// 	return numA / numB, nil
// }

func (s exchangeService) ListExchange(ctx context.Context) ([]ExchangeType, error) {
	fmt.Println("hueheuhedasdasdu")
	var a []ExchangeType
	return a, nil
}

// loadFeatures loads features from a JSON file.
func (s *exchangeService) loadFeatures(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			otherLog.Fatalf("Failed to load default features: %v", err)
		}
	} else {
		data = exchangesData
	}
	if err := json.Unmarshal(data, &s.savedFeatures); err != nil {
		otherLog.Fatalf("Failed to load default features: %v", err)
	}
}

var exchangesData = []byte(`[
{
	"exchange_id": "OKCOIN_CNY",
    "website": "https://www.okcoin.cn/",
    "name": "OKCoin CNY",
	"upvotes": 3,
	"downvotes": "8"
},
{
	"exchange_id": "HUOBI",
	"website": "https://www.huobi.com/",
	"name": "Huobi (HBUS)",
	"upvotes": 10,
	"downvotes": "5"
},
{
    "exchange_id": "OVEX",
    "website": "https://www.ovex.io/",
    "name": "OVEX",
	"upvotes": 12,
	"downvotes": "9"
}
]`)
