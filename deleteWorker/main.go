package main

import (
	"delete/Global"
	"delete/RabbitMq"
	"log"
)

func main() {
	RabbitMq.InitRabbitMq()
	defer Global.Ch.Close()
	defer Global.Conn.Close()
	chan1 := make(chan int)
	log.Println("Delete Worker Running!")
	<-chan1

}
