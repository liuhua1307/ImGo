package configs

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"sync"
)

var once *sync.Once
var (
	newProducer rocketmq.Producer
	err         error
)

func NewProducer() rocketmq.Producer {
	once.Do(func() {
		newProducer, err = rocketmq.NewProducer(
			producer.WithNameServer([]string{os.Getenv("ROCKETMQ_ADDR")}),
			producer.WithRetry(2),
			producer.WithGroupName("GID_Dynamic"),
			producer.WithNamespace("NS_Dynamic"),
		)
		if err != nil {
			panic(err)
		}
	})
	return newProducer
}
