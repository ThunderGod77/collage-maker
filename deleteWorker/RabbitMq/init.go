package RabbitMq

import (
	"delete/Del"
	"delete/Global"
	"github.com/streadway/amqp"
	"log"
)

var err error

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func InitRabbitMq() {
	Global.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	Global.Ch, err = Global.Conn.Channel()
	failOnError(err, "Failed to open a channel")

	_, err = Global.Ch.QueueDeclare(
		"deleteWorker", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgsW, err := Global.Ch.Consume(
		"deleteWorker", // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")

	go Del.WReceive(msgsW)

}
