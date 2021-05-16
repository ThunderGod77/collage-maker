package PublishResult

import (
	"collageWorker/Global"
	"github.com/streadway/amqp"
	"log"
)

func wPublish(message []byte, folderName string) {

	err := Global.Ch.Publish(
		"",                  // exchange
		"receiveFromWorker", // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	//failOnError(err, "Failed to publish a message for folder "+folderName)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Work done on " + folderName)
}
