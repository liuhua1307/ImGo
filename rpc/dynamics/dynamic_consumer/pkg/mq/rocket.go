package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
)

type Rocket struct {
	c rocketmq.PullConsumer
}

func NewRocket(c rocketmq.PullConsumer) *Rocket {
	return &Rocket{c: c}
}

func NewMq(r *Rocket) Mq {
	return r
}

func (r *Rocket) Consume(ctx context.Context) error {
	//todo consume message
	return nil
}
