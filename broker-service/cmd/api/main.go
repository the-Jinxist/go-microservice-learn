package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const WEBPORT = "8080"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {

	//connect to rabbit mq
	conn, err := connect()
	if err != nil {
		log.Fatalf("cannot connect to rabbitmq, %s", err)
	}

	defer conn.Close()

	app := Config{
		Rabbit: conn,
	}

	log.Printf("Starting broker service on port %s", WEBPORT)

	//define http server
	serve := http.Server{
		Addr:    fmt.Sprintf(":%s", WEBPORT),
		Handler: app.routes(),
	}

	//start the server
	err = serve.ListenAndServe()
	if err != nil {
		log.Fatalf("error while listening and serving: %s", err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var connection *amqp.Connection

	backoff := 1 * time.Second

	//don't continue until rabbit is ready

	//we updated the rabbitmq link from  amqp://guest:guest@localhost to amqp://guest:guest@rabbitmq -- the rabbitmq at the end matches the name of the docker service in docker-compose.yml
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("rabbitmq not yet ready")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			log.Printf("Niggling rabbitmq error: %s", err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off")

		time.Sleep(backoff)
		continue
	}

	return connection, nil

}
