package main

import (
	"context"
	"fmt"
	"log"
	"softwareIIbackend/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// initialize config
	config := config.New()
	cfg := config.Database

	// connect to mongodb
	ctx := context.Background()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&appName=Cluster0", cfg.User, cfg.Password, cfg.Host, cfg.DBName)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	log.Println("Successful connection with MongoDB")

	err = SeedUsers(ctx, client.Database(cfg.DBName).Collection("users"))
	if err != nil {
		panic(err)
	}

	log.Println("Successful db migration")
}
