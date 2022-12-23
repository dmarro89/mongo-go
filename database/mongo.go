package database

import (
	"context"
	"errors"
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

func NewClient(client *mongo.Client) *MongoDB {
	mongoDB := MongoDB{client: client}
	return &mongoDB
}

func (mongoDB *MongoDB) getConnectionURI() string {
	return os.Getenv(connectionURI)
}

func (mongoDB *MongoDB) Connect(ctx context.Context) error {
	var err error
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(mongoDB.getConnectionURI()).
		SetServerAPIOptions(serverAPIOptions)
	mongoDB.client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	return nil
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
