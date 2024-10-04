package cache

import (
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/domain"
)

var ProviderSet = wire.NewSet(NewDynamicsRepository, wire.Bind(new(domain.DynamicRepository), new(*DynamicsRedisRepository)))
