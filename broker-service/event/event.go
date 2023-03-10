package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name of the exchange
		"topic",      //type
		true,         //is the exchange durable
		false,        //don't delete the event when we're done with it,
		false,        //false because we're using it between microservices
		false,        //no wait
		nil,          //no specific arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //name
		false, //durable, false, get rid of it when we're done with it
		false, //delete when unused
		true,  //yes, this is an exclusive channel,
		false, //not noWait,
		nil,   //no other arguments for now
	)
}
