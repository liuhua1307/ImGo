package cache

import (
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/domain"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/pkg/cache"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/pkg/mq"
)

type DynamicRedisRepository struct {
	cache cache.Cache
	mq    mq.Mq
}

func NewDynamicRepository(cache cache.Cache, mq mq.Mq) *DynamicRedisRepository {
	return &DynamicRedisRepository{cache: cache, mq: mq}
}

func (d DynamicRedisRepository) PushDynamic(info domain.Dynamic) error {
	//todo push dynamic
}
