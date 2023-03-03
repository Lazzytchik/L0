package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	nc, connectionError := nats.Connect(nats.DefaultURL)

	if connectionError != nil {
		connectionError.Error()
	}

	js, jetError := nc.JetStream()

	if jetError != nil {
		jetError.Error()
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "FOO",
		Subjects: []string{"foo"},
	})

	js.Publish("foo", []byte("Hello JS!"))

	// Publish messages asynchronously.
	for i := 0; ; i++ {
		msg := fmt.Sprintf("Hello JS Async %d!", i)
		js.PublishAsync("foo", []byte(msg))
		log.Printf("Sent: %s \n", msg)

		time.Sleep(2 * time.Second)
	}

}
