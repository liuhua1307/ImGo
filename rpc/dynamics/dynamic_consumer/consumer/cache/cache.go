package cache

import (
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/domain"
)

var ProviderSet = wire.NewSet(NewDynamicRepository, wire.Bind(new(domain.DynamicRepository), new(*DynamicRedisRepository)))
