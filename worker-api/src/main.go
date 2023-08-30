package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// load configuration
	loadConfig()

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUser, rabbitPass, config.RabbitMQ.Host, config.RabbitMQ.Port)

	// Create a new RabbitMQ connection.
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Queue declaration
	q, err := ch.QueueDeclare(
		config.RabbitMQ.Queue.Name, // name
		true,                       // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// QoS configuration
	err = ch.Qos(
		config.RabbitMQ.QoS.PrefetchCount, // prefetch count
		config.RabbitMQ.QoS.PrefetchSize,  //prefetch size
		config.RabbitMQ.QoS.Global,        //global
	)
	failOnError(err, "Failed to configure Qos")

	// Register consumer
	msgs, err := ch.Consume(
		q.Name,                        // queue
		config.RabbitMQ.Consumer.Name, // consumer
		false,                         // auto-ack
		false,                         // exclusive
		false,                         // no-local
		false,                         // no-wait
		nil,                           // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			// process message
			process(d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
