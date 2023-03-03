package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
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

	//wg := sync.WaitGroup{}
	//wg.Add(5)

	go js.Subscribe("foo", func(msg *nats.Msg) {
		log.Printf("Got: %s", string(msg.Data))
		//wg.Done()
	}, nats.DeliverNew())

	http.ListenAndServe(":3333", nil)
	//wg.Wait()

	log.Println("THE END")

}
