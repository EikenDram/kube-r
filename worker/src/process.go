package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func process(message amqp.Delivery) {
	//
	log.Printf(" > Received message: %s\n", message.Body)
}
