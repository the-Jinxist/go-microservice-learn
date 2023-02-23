package main

import (
	"fmt"
	"log"
	"net/http"
)

const WEBPORT = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s", WEBPORT)

	//define http server
	serve := http.Server{
		Addr:    fmt.Sprintf(":%s", WEBPORT),
		Handler: app.routes(),
	}

	//start the server
	err := serve.ListenAndServe()
	if err != nil {
		log.Fatalf("error while listening and serving: %s", err)
	}
}
