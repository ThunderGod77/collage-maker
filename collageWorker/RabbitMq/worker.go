package RabbitMq

import (
	"bytes"
	"collageWorker/Global"
	"collageWorker/ImageManipulation"
	"encoding/gob"
	"log"

	"github.com/streadway/amqp"
)



func wReceive(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message(by worker)")
		var imgInfo Global.ImagesInfo
		data := bytes.NewBuffer(d.Body)
		dec := gob.NewDecoder(data)
		err = dec.Decode(&imgInfo)
		if err != nil {
			log.Println(err)

			return
		}

		go ImageManipulation.ImagePreProcessing(imgInfo)

	}
}
