package server

import (
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/kitex_gen/api"
	consumerApi "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/kitex_gen/api/consumer"
)

var ProviderSet = wire.NewSet(NewConsumerImpl, NewKitexServer)

func NewKitexServer(consumer api.Consumer) server.Server {
	return consumerApi.NewServer(consumer)
}
