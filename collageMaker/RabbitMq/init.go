package RabbitMq

import (
	"github.com/streadway/amqp"
	"log"
)

var Ch *amqp.Channel
var Conn *amqp.Connection
var err error

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func InitRabbitMq() {
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	Ch, err = Conn.Channel()
	failOnError(err, "Failed to open a channel")

	_, err = Ch.QueueDeclare(
		"sendToWorker", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")
	_, err = Ch.QueueDeclare(
		"receiveFromWorker", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")
	_, err = Ch.QueueDeclare(
		"deleteWorker", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := Ch.Consume(
		"receiveFromWorker", // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	failOnError(err, "Failed to register a consumer")

	go Receive(msgs)

}
