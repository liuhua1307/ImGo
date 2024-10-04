package server

import (
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/kitex_gen/api"
	spaceApi "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/kitex_gen/api/space"
)

var ProviderSet = wire.NewSet(NewSpaceImpl, NewKitexServer)

func NewKitexServer(space api.Space) server.Server {
	return spaceApi.NewServer(space)
}
