package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer nc.Close()

	var sub = "test.subject"
	var que = "queue.foo"

	nc.QueueSubscribe(sub, que, func(msg *nats.Msg) {
		log.Println("Subscriber 1:", string(msg.Data))
	})

	nc.QueueSubscribe(sub, que, func(msg *nats.Msg) {
		log.Println("Subscriber 2:", string(msg.Data))
	})

	nc.QueueSubscribe(sub, que, func(msg *nats.Msg) {
		log.Println("Subscriber 3:", string(msg.Data))
	})

	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Message %d", i)

		if err := nc.Publish(sub, []byte(message)); err != nil {
			log.Fatalln(err)
		}
	}

	time.Sleep(2 * time.Second)
}
