package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type sendMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload sendMessage
	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		log.Printf("error with sending mail: %s, request payload: %v", err, requestPayload)
		app.errorJson(w, err)
		return
	}

	var msg Message
	msg.From = requestPayload.To
	msg.To = requestPayload.To
	msg.Subject = requestPayload.Subject
	msg.Data = requestPayload.Message

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Printf("error with sending mail: %s, message: %v", err, msg)
		app.errorJson(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Sent to " + requestPayload.To

	app.writeJson(w, http.StatusAccepted, payload)
}
