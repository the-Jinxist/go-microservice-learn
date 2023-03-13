package main

import (
	"context"
	"log"
	"logger-service/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

// Every method that you want to be accessed in RPC server must be exported(must start with a capital letter)
func (r *RPCServer) LogInfo(payload RPCPayload, response *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*response = "Processed payload via RPC" + payload.Name
	return nil
}
