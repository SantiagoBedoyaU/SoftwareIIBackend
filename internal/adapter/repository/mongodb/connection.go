package mongodb

import (
	"context"
	"fmt"
	"log"
	"softwareIIbackend/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection struct {
	DBName string
	Client *mongo.Client
}

func NewMongodbConnection(ctx context.Context, cfg config.DatabaseConfig) (*MongoDBConnection, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&appName=Cluster0", cfg.User, cfg.Password, cfg.Host, cfg.DBName)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	log.Println("Successfull connection with MongoDB")
	return &MongoDBConnection{
		DBName: cfg.DBName,
		Client: client,
	}, nil
}

func (m *MongoDBConnection) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

func (m *MongoDBConnection) GetDatabase() *mongo.Database {
	return m.Client.Database(m.DBName)
}
