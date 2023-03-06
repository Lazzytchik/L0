package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"lazzytchik/L0/models"
	nats2 "lazzytchik/L0/nats"
	"lazzytchik/L0/storage"
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

	orderStorage := storage.New()

	js.Core.Subscribe("foo", func(msg *nats.Msg) {
		log.Printf("Got: %s", string(msg.Data))
		order := models.Order{}
		json.Unmarshal(msg.Data, &order)
		orderStorage.Add(order)
	}, nats.DeliverNew())

	time.Sleep(10 * time.Second)

	//http.ListenAndServe(":3333", nil)

	log.Println("THE END")

}
