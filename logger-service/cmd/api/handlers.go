package main

import (
	"logger-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {

	var requestPayload JSONPayload
	_ = app.readJson(w, r, requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Event logged!",
	}

	app.writeJson(w, http.StatusAccepted, response)

}