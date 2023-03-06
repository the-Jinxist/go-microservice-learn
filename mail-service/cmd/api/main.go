package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const WEBPORT = "80"

func main() {
	app := Config{}

	log.Printf("Starting mail service on port %s", WEBPORT)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", WEBPORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error while listening and serving: %s", err)
	}
}
