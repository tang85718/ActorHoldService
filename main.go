package main

import (
	"time"
	"github.com/streadway/amqp"
	"fmt"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	failOnError(err)

	ch, err := conn.Channel()
	failOnError(err)
	defer ch.Close()

	//q, err := ch.QueueDeclare(
	//	"actors",
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//failOnError(err)

	//ms, err := mgo.Dial("")
	//failOnError(err)

	//actorCol := ms.DB("crm").C("actors")

	for {
		fmt.Println("actor_hold_service running..")
		time.Sleep(time.Second)
	}

}
