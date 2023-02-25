package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// We're using a raw http handler, my chest.
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")

	//Setting the header of response to accept json and the response code to 202
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	//Actually sending out the payload
	w.Write(out)
}
