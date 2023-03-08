package stream

import (
	"github.com/nats-io/nats.go"
)

type JetsStream struct {
	Core nats.JetStreamContext
}

func (js *JetsStream) Connect() {
	nc, connectionError := nats.Connect(nats.DefaultURL)

	if connectionError != nil {
		connectionError.Error()
	}

	var jetError error
	js.Core, jetError = nc.JetStream()

	if jetError != nil {
		jetError.Error()
	}
}

func (js *JetsStream) AddStream(name string) {
	js.Core.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{"order"},
	})
}

func (js *JetsStream) SubscribeAsync(handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	return js.Core.Subscribe("order", handler, nats.DeliverNew())

}
