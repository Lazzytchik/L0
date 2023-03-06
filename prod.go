package main

import (
	"github.com/nats-io/nats.go"
	"lazzytchik/L0/json"
	nats2 "lazzytchik/L0/nats"
	"log"
	"time"
)

func main() {

	js := nats2.JetsStream{}
	js.Connect()

	js.Core.AddStream(&nats.StreamConfig{
		Name:     "FOO",
		Subjects: []string{"foo"},
	})

	order := json.Order{}

	// Publish messages asynchronously.
	for i := 0; ; i++ {
		msg := order.Generate()
		js.Core.PublishAsync("foo", msg)
		log.Printf("Sent: %s \n", msg)

		time.Sleep(2 * time.Second)
	}

}
