package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/google/wire"
)

var _ Cache = &RedisCache{}

type RedisCache struct {
	pool *redis.Client
}

func (r *RedisCache) ZSet(key string, z []*redis.Z, duration time.Duration) error {
	err := r.pool.ZAdd(r.pool.Context(), key, z...).Err()
	if err != nil {
		return err
	}
	r.pool.ExpireAt(context.Background(), key, time.Now().Add(duration))
	return nil
}
func (r *RedisCache) ZCard(key string) (int64, error) {
	return r.pool.ZCard(context.Background(), key).Result()

}

func (r *RedisCache) ZRange(key string, start, end int64) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) ZRevRange(key string, start, end int64) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) ZRangeUnMarshal(key string, start, end int64, value []interface{}) error {
	result, err := r.pool.ZRange(r.pool.Context(), key, start, end).Result()
	if err != nil {
		return err
	}
	for _, v := range result {
		err := json.Unmarshal([]byte(v), &value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisCache) ZRevRangeUnMarshal(key string, start, end int64, value []interface{}) error {
	//TODO implement me
	panic("implement me")
}

var ProviderSet = wire.NewSet(NewRedis, wire.Bind(new(Cache), new(*RedisCache)))

func NewRedis(pool *redis.Client) *RedisCache {
	return &RedisCache{
		pool: pool,
	}
}

func (r *RedisCache) Set(key string, value string, duration time.Duration) error {
	return r.pool.Set(r.pool.Context(), key, value, duration).Err()
}

func (r *RedisCache) SetMarshal(key string, value interface{}, duration time.Duration) error {
	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Set(key, string(byteValue), duration)
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.pool.Get(r.pool.Context(), key).Result()
}

func (r *RedisCache) GetUnMarshal(key string, value interface{}) error {
	res, err := r.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(res), value)
}

func (r *RedisCache) Del(keys ...string) error {
	var err error
	for _, key := range keys {
		if err = r.pool.Del(r.pool.Context(), key).Err(); err != nil && err != redis.Nil {
			return err
		}
	}
	return nil
}
