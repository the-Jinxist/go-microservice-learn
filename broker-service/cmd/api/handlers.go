package main

import (
	"net/http"
)

// We're using a raw http handler, my chest.
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}

	_ = app.writeJson(w, http.StatusOK, payload)
}
