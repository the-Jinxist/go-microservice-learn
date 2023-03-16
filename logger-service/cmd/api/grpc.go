package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"logger-service/logs"
	"net"
	"time"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, logRequest *logs.LogRequest) (*logs.LogResponse, error) {
	input := logRequest.GetLogEntry()

	//write the log
	logEntry := data.LogEntry{
		Name:      input.Name,
		Data:      input.Data,
		CreatedAt: time.Now(),
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{
			Result: "Request failed!",
		}
		return res, err
	}

	res := &logs.LogResponse{
		Result: "Result logged!",
	}
	return res, nil
}

func (app *Config) grpcListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for grpc: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{
		Models: app.Models,
	})

	log.Printf("grpc server started on port: %s", gRpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for grpc: %v", err)
	}
}
