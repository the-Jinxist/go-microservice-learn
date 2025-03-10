package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"net"
	"net/http"
	"net/rpc"
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
	Models data.Models
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

	app := Config{
		Models: data.New(client),
	}

	//We have to register the rpc server first
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()

	//Starting and listening to the grpcServer
	go app.grpcListen()

	app.serve()

}

func (app *Config) rpcListen() error {
	log.Println("starting rpc server on port", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}

	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(rpcConn)
	}

	return nil
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error ListenAndServe(): %s", err)
	}
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
