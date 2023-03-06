package nats

import (
	"github.com/nats-io/nats.go"
)

type BaseSubscriber struct {
	JS      nats.JetStreamContext
	Channel string
}
