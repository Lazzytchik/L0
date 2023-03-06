package nats

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
