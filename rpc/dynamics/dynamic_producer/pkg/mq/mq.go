package mq

import "context"

type Mq interface {
	// Publish a message to the queue
	Publish(ctx context.Context, message ...interface{}) error
	// Consume a message from the queue
}
