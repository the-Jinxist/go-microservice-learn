package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {

	//connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Fatalf("Error while connecting to mongo: %s", err)
	}

	client = mongoClient

	//create a client in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	//close mongo connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("error with disconnecting: %s", err)
		}

	}()

}

func connectToMongo() (*mongo.Client, error) {
	//create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return c, nil
}
