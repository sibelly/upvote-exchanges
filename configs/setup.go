package configs

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

const (
	DB         = "upvotes"
	COLLECTION = "exchanges"
)

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(GetEnvVar("MONGOURI"))
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		fmt.Println("Connected to MongoDB")
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

//getting database collections
func GetCollection(collectionName *string) (*mongo.Collection, error) {
	// Get singleton client
	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DB).Collection(COLLECTION)
	if collectionName != nil {
		collection = client.Database(DB).Collection(*collectionName)
	}

	return collection, nil
}
