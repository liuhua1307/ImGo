package configs

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"sync"
)

var (
	once sync.Once
	c    rocketmq.PullConsumer
	err  error
)

func NewConsumer() rocketmq.PullConsumer {
	once.Do(func() {
		c, err = rocketmq.NewPullConsumer(
			consumer.WithGroupName("GID_Dynamic"),
			consumer.WithNameServer([]string{os.Getenv("ROCKETMQ_ADDR")}),
			consumer.WithNamespace("NS_Dynamic"),
			consumer.WithCredentials(primitive.Credentials{
				AccessKey: os.Getenv("ROCKETMQ_ACCESS_KEY"),
				SecretKey: os.Getenv("ROCKETMQ_SECRET_KEY"),
			}),
		)
	})
	return c
}
