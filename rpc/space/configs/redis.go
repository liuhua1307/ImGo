package configs

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/liuhua1307/ImGo/rpc/space/pkg/log"
	"github.com/spf13/viper"
	"sync"
)

var (
	rdb     *redis.Client
	onceRdb sync.Once
)

func NewRDB(ctx context.Context) *redis.Client {
	onceRdb.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:         viper.GetString("redis.addr"),
			Password:     viper.GetString("REDIS_PASSWORD"),
			DB:           viper.GetInt("redis.db"),
			PoolSize:     viper.GetInt("redis.pool_size"),
			MinIdleConns: viper.GetInt("redis.min_idle_conn"),
		})

		if _, err := rdb.Ping(ctx).Result(); err != nil {
			panic(err)
		}
	})
	return rdb.WithContext(ctx)
}

func CloseRedis() {
	if err := rdb.Close(); err != nil {
		log.Log().Info("Close redis failed", log.Field{Key: "error", Value: err})
	}
}
