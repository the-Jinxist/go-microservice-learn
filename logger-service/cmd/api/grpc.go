package main

import (
	"context"
	"logger-service/data"
	"logger-service/logs"
	"time"
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
