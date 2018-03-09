package main

import (
	"time"
	"github.com/streadway/amqp"
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/tangxuyao/mongo"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	conn, err := amqp.Dial("amqp://127.0.0.1:5672/")
	failOnError(err)

	ch, err := conn.Channel()
	failOnError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"actors",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err)

	ms, err := mgo.Dial("")
	failOnError(err)

	colActors := ms.DB("crm").C("actors")

	for {
		var results []mongo.Charactor
		colActors.Find(nil).All(&results)

		if len(results) > 0 {


		}


		fmt.Println("actor_hold_service running..")
		time.Sleep(time.Second)
	}

}
