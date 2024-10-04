// Package cache ，定义了 cache 相关接口，包括 Set、Get、Del 等方法。在不同的模块中的 repository 层，可以通过 cache 接口来实现缓存。具体可以参考 /internal/task/cache/task_redis.go 中的实现。
// 参考 go-clen-arch 中的 issue，使用缓存相关的通过实现对应模块的 repository 层的接口，调用对应 database repository 的方法实现，而不是通过在 usecase 层实现。
package cache

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache interface {
	Set(key string, value string, duration time.Duration) error
	SetMarshal(key string, value interface{}, duration time.Duration) error
	ZSet(key string, zs []*redis.Z, duration time.Duration) error
	ZRange(key string, start, end int64) ([]string, error)
	ZRevRange(key string, start, end int64) ([]string, error)
	ZRangeUnMarshal(key string, start, end int64, values []interface{}) error
	ZRevRangeUnMarshal(key string, start, end int64, values []interface{}) error
	ZCard(key string) (int64, error)
	Get(key string) (string, error)
	GetUnMarshal(key string, value interface{}) error
	Del(keys ...string) error
}
