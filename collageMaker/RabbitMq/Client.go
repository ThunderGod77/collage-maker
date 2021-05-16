package RabbitMq

import (
	"bytes"
	"collageMaker/Global"
	"encoding/gob"
	"github.com/streadway/amqp"
	"log"
)

func Publish(message []byte, folderId string) {
	body := message
	err := Ch.Publish(
		"",             // exchange
		"sendToWorker", // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message to worker queue")
	if err != nil {
		delete(Global.TaskBeingDone, folderId)
		Global.TaskError[folderId] = true
	}

}

func Receive(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message(by maker)")
		var workInfo Global.WorkCompletion
		data := bytes.NewBuffer(d.Body)
		dec := gob.NewDecoder(data)
		err = dec.Decode(&workInfo)
		if err != nil {
			log.Println(err)

			return
		}
		if workInfo.Err {
			log.Println("Error making collage " + workInfo.FolderId)
			delete(Global.TaskBeingDone, workInfo.FolderId)
			Global.TaskError[workInfo.FolderId] = true

		} else {
			log.Println("Work done of folder " + workInfo.FolderId)
			delete(Global.TaskBeingDone, workInfo.FolderId)
			Global.TaskCompleted[workInfo.FolderId] = true
		}

	}
}
func PublishDelete(message []byte) {
	body := message
	err := Ch.Publish(
		"",             // exchange
		"deleteWorker", // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message to delete queue")


}