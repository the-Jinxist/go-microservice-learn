package main

import (
	"listener/event"
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

	//start listening for messages
	log.Println("listening for messages")

	//create consumer
	consumer, err := event.NewConsumer(conn)
	if err != nil {
		log.Fatalf("cannot get new consumer, %s", err)
	}

	//watch and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Printf("error while listening to rabbitmq, %s", err)
	}

	log.Println("connected to rabbitmq")

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
