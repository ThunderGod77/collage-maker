package main

import (
	"collageWorker/Global"
	"collageWorker/RabbitMq"
	"log"
)

func main() {

	RabbitMq.InitRabbitMq()
	defer Global.Conn.Close()
	defer Global.Ch.Close()
	chan1 := make(chan int)
	log.Println("Worker Running!")
	<-chan1
}
