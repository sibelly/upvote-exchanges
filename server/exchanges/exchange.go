package exchanges

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/sibelly/upvote-exchanges/configs"
	"github.com/sibelly/upvote-exchanges/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ExchangeType represents the bson readable data from the protobuf
type ExchangeType struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	ExchangeId string             `bson:"exchange_id"`
	Name       string             `bson:"name"`
	Website    string             `bson:"website"`
	Upvotes    int32              `bson:"upvotes"`
	Downvotes  int32              `bson:"downvotes"`
}

// ExchangeServiceServer is a implementation of CryptoService provided by gRPC
type ExchangeServiceServer struct {
	*pb.UnimplementedExchangesServiceServer
	savedFeatures []interface{} // read-only after initialized
}

// NewExchangeServiceServer returns a pointer to a ExchangeServiceServer
func NewExchangeServiceServer() *ExchangeServiceServer {
	return &ExchangeServiceServer{}
}

// Upvote takes a VoteRequest and updates the "upvotes" field on a given crypto
// returning a VoteResponse if successful.
func (s *ExchangeServiceServer) Upvote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteResponse, error) {
	log.Printf("Received Upvote request for %v", in.GetExchangeId())
	readExchange := ExchangeType{}

	// Load collection
	collectionName := "exchanges"
	collection, err := configs.GetCollection(&collectionName)
	if err != nil {
		return nil, err
	}

	// Read exchange of a given name
	res := collection.FindOne(context.TODO(), bson.M{"exchange_id": in.GetExchangeId()})
	if err := res.Decode(&readExchange); err != nil {
		return nil, status.Error(codes.NotFound, "Could not find Object")
	}

	// Update object
	readExchange.Upvotes += 1
	_, err = collection.ReplaceOne(context.TODO(), primitive.M{"exchange_id": readExchange.ExchangeId}, readExchange)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update exchange: %v", err)
	}

	return &pb.VoteResponse{
		Name:    readExchange.Name,
		Website: readExchange.Website,
		Votes:   readExchange.Upvotes - readExchange.Downvotes,
	}, nil
}

// ListExchange takes an Empty request, returning a stream of Exchange
func (s *ExchangeServiceServer) ListExchanges(in *pb.Empty, stream pb.ExchangesService_ListExchangeServer) error {
	log.Print("Received List request")
	// Load collection
	collection, err := configs.GetCollection(nil)
	if err != nil {
		return err
	}
	// Create Mongo cursor
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, "Error: %v", err)
	}
	defer cursor.Close(context.TODO())
	// Iterate over cursor entries
	for cursor.Next(context.TODO()) {
		item := &ExchangeType{}
		if err := cursor.Decode(item); err != nil {
			return status.Errorf(codes.Internal, "Could not decode data: %v", err)
		}
		stream.Send(&pb.Exchange{
			Name:      item.Name,
			Website:   item.Website,
			Upvotes:   item.Upvotes,
			Downvotes: item.Downvotes,
		})
	}
	// Check for cursor errors
	if err = cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "Error: %v", err)
	}
	return nil
}

// loadFeatures loads features from a JSON file.
func (s *ExchangeServiceServer) LoadFeatures(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to load default features: %v", err)
		}
	} else {
		data = exchangesData
	}
	// Load collection
	collection, err := configs.GetCollection(nil)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &s.savedFeatures); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	result, err := collection.InsertMany(context.TODO(), s.savedFeatures)
	if err != nil {
		panic(err)
	}
	log.Println("Result initializing => ", result)
}

var exchangesData = []byte(`[
{
	"exchange_id": "OKCOIN_CNY",
    "website": "https://www.okcoin.cn/",
    "name": "OKCoin CNY",
	"upvotes": 3,
	"downvotes": 8
},
{
	"exchange_id": "HUOBI",
	"website": "https://www.huobi.com/",
	"name": "Huobi (HBUS)",
	"upvotes": 10,
	"downvotes": 5
},
{
    "exchange_id": "OVEX",
    "website": "https://www.ovex.io/",
    "name": "OVEX",
	"upvotes": 12,
	"downvotes": 9
}
]`)
