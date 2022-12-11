package database

import (
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionURI        = "CONNECTION_URI"
	ClientNotInitialized = "client not initialized"
	Database             = "mongo-go"
)

var ErrClientNotInitialized = errors.New(ClientNotInitialized)

type MongoDB struct {
	client *mongo.Client
}

func New() *MongoDB {
	mongoDB := MongoDB{}
	return &mongoDB
}

func (mongoDB *MongoDB) GetConnectionURI() string {
	return os.Getenv(connectionURI)
}

func (mongoDB *MongoDB) Connect() error {
	var err error
	ctx := context.Background()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(mongoDB.GetConnectionURI()).
		SetServerAPIOptions(serverAPIOptions)
	mongoDB.client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (mongoDB *MongoDB) Disconnect(ctx context.Context) error {
	if mongoDB.client == nil {
		return ErrClientNotInitialized
	}
	return mongoDB.client.Disconnect(ctx)
}

func (mongoDB *MongoDB) GetClient() *mongo.Client {
	return mongoDB.client
}

func (mongoDB *MongoDB) GetCollection(database, collection string) *mongo.Collection {
	return mongoDB.client.Database(database).Collection(collection)
}
