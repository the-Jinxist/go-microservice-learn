package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	//validate user against database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	matches, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !matches {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJson(w, http.StatusAccepted, payload)

}
