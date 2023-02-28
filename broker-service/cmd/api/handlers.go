package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// We're using a raw http handler, my chest.
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}

	_ = app.writeJson(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJson(w, r, requestPayload)
	if err != nil {
		app.errorJson(w, err)
	}

	switch requestPayload.Action {
	case "auth":
		{
			app.authenticate(w, requestPayload.Auth)

		}
	default:
		app.errorJson(w, errors.New("unknown action"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	//create some json for sending to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	//call the service
	//to call the auth microservie, we make the equivalent of an http request, to access the service we use the name of service that we
	//added in our docker-compose.yml

	//the authentication microservice is written as `authentication-service` on line 36 in the docker-compose.yml
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	defer response.Body.Close()

	//get back the correct status code

	if response.StatusCode == http.StatusBadRequest {
		app.errorJson(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJson(w, errors.New("error calling auth service"))
		return
	}

	var jsonFromServiceResponse jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromServiceResponse)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	if jsonFromServiceResponse.Error {
		app.errorJson(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromServiceResponse

	app.writeJson(w, http.StatusAccepted, payload)
}
