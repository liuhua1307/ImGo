package cache

import (
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/space/domain"
)

var ProviderSet = wire.NewSet(NewDynamicsRepository, wire.Bind(new(domain.DynamicRepository), new(*DynamicsRedisRepository)))
