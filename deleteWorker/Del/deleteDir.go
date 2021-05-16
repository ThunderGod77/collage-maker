package Del

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func DeleteDirectory(folderId string) {
	channel := make(chan string, 2)
	select {
	case output := <-channel:
		fmt.Println(output)
	case <-time.After(15*60* time.Second):
		err := os.RemoveAll("./../" + folderId)
		if err != nil {
			log.Println("Error removing folder " + folderId)
		}
		log.Println("Removed folder " + folderId)
	}
}
func WReceive(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message(by delete worker)")
		DeleteDirectory(string(d.Body))

	}
}
