package main

import (
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

	sub, _ := js.SubscribeSync("foo", nats.DeliverNew())

	for {
		msg, msgErr := sub.NextMsg(time.Second)

		if msgErr != nil {
			log.Println(msgErr)
			break
		}

		log.Println(string(msg.Data))
	}

	log.Println("THE END")

	//js.AddStream(&nats.StreamConfig{
	//	Name:     "FOO",
	//	Subjects: []string{"FOO"},
	//	MaxBytes: 1024,
	//})
	//
	//_, consumerError := js.AddConsumer("FOO", &nats.ConsumerConfig{
	//	Durable: "BAR",
	//})
	//
	//if consumerError != nil {
	//	consumerError.Error()
	//}
	//
	//info, _ := js.StreamInfo("FOO")
	//
	//log.Printf("Channel \"%s\" created on %s \n", info.Config.Name, info.Created)
	//
	//_, pubError := js.Publish("FOO", []byte("Hello Jet!"))
	//
	//if pubError != nil {
	//	pubError.Error()
	//}
	//
	//_, subErr := js.Subscribe("FOO", func(msg *nats.Msg) {
	//	meta, _ := msg.Metadata()
	//	log.Printf("Stream sequence:   %v\n", meta.Sequence.Stream)
	//	log.Printf("Consumer sequence: %v\n", meta.Sequence.Consumer)
	//}, nats.OrderedConsumer())
	//
	//if subErr != nil {
	//	subErr.Error()
	//}

}
