package main

import (
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//connect to rabbit mq
	conn, err := connect()
	if err != nil {
		log.Fatalf("cannot connect to rabbitmq, %s", err)

	}

	defer conn.Close()
	log.Println("connected to rabbitmq")

	//start listening for messages

	//create consumer

	//watch and consume events
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var connection *amqp.Connection

	backoff := 1 * time.Second

	//don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
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
