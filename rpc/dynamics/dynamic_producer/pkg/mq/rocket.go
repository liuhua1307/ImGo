package mq

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/google/wire"
	"sync"
)

const Topic = "Dynamic"

var ProviderSet = wire.NewSet(NewRocket, NewMQ)

type Rocket struct {
	producer rocketmq.Producer
}

func NewRocket(producer rocketmq.Producer) *Rocket {
	return &Rocket{producer: producer}
}

func NewMQ(r *Rocket) Mq {
	return r
}

func (r *Rocket) Publish(ctx context.Context, message ...interface{}) error {
	var (
		n       = len(message)
		wg      sync.WaitGroup
		errChan = make(chan error, n)
	)

	if n == 0 {
		return errors.New("message is empty")
	}

	wg.Add(n)

	for _, m := range message {
		go func(m interface{}) {
			defer wg.Done()
			//todo publish message

			bytes, err := json.Marshal(m)
			if err != nil {
				errChan <- err
				return
			}

			err = r.producer.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, e error) {
				if e != nil {
					errChan <- e
					return
				}
			}, primitive.NewMessage(Topic, bytes))
			if err != nil {
				errChan <- err
				return
			}
		}(m)

		// check errChan is Empty, if errChan is not empty, return error, else break and return nil
		select {
		case err := <-errChan:
			return err
		default:
			break
		}
	}

	wg.Wait()
	return nil
}
