package mq

import (
	"context"
	"github.com/google/wire"
)

type Mq interface {
	// Consume a message from the queue
	Consume(ctx context.Context) error
}

var ProviderSet = wire.NewSet(NewRocket, NewMq)
