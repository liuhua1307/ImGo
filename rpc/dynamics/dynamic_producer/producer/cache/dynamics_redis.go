package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/domain"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/pkg/cache"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/pkg/mq"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/producer/repository"
	"strconv"
	"time"
)

const UserTimeLinePrefix = "UserTimeLine-"

type DynamicsRedisRepository struct {
	dynamicsDBRepository domain.DynamicRepository
	pool                 cache.Cache
	mq                   mq.Mq
}

func NewDynamicsRepository(pool cache.Cache, dbRepo *repository.DynamicsMYSQLRepository, m mq.Mq) *DynamicsRedisRepository {
	return &DynamicsRedisRepository{pool: pool, dynamicsDBRepository: dbRepo, mq: m}
}

func (h DynamicsRedisRepository) GetFocusDynamic(ctx context.Context, targetId, limit, userId int64) ([]domain.Dynamics, error) {
	empty, err := h.checkZSetNotEmpty(userId)
	if empty {
		return h.getFocusDynamicWithPull(ctx, targetId, limit, userId)
	}
	if err != nil {
		return nil, err
	}
	var results []interface{}
	err = h.pool.ZRangeUnMarshal(UserTimeLinePrefix+strconv.FormatInt(userId, 10), 0, -1, results)
	if err != nil {
		return nil, err
	}
	var dynamics []domain.Dynamics
	for _, result := range results {
		dynamics = append(dynamics, result.(domain.Dynamics))
	}
	return dynamics, nil
}

func (h DynamicsRedisRepository) PushDynamicAndPush(ctx context.Context, info domain.DynamicInfo) error {
	err := h.dynamicsDBRepository.PushDynamicAndPush(ctx, info)
	if err != nil {
		return err
	}
	list, err := h.dynamicsDBRepository.GetFocusList(ctx, int64(info.AuthorId))
	if err != nil {
		return err
	}
	err = h.mq.Publish(ctx, info, list)
	if err != nil {
		return err
	}
	return nil
}

func (h DynamicsRedisRepository) getFocusDynamicWithPull(ctx context.Context, targetId, limit, userId int64) ([]domain.Dynamics, error) {
	dynamics, err := h.dynamicsDBRepository.GetFocusDynamic(ctx, targetId, limit, 0)
	if err != nil {
		return nil, err
	}
	var redisZ []*redis.Z
	for _, v := range dynamics {
		score := float64(v.CreatedAt.Unix()) + (float64(v.ID) / float64(10^len(strconv.FormatInt(int64(v.ID), 10))))
		redisZ = append(redisZ, &redis.Z{
			Score:  score,
			Member: v,
		})
	}
	err = h.pool.ZSet(UserTimeLinePrefix+strconv.FormatInt(userId, 10), redisZ, time.Hour*72)
	if err != nil {
		return nil, err
	}
	return dynamics, nil
}

func (h DynamicsRedisRepository) checkZSetNotEmpty(userId int64) (bool, error) {
	card, err := h.pool.ZCard(UserTimeLinePrefix + strconv.FormatInt(userId, 10))
	return errors.Is(err, redis.Nil) || (err == nil && card <= 1), err
}

func (h DynamicsRedisRepository) GetFocusList(ctx context.Context, userId int64) ([]int64, error) {
	return h.dynamicsDBRepository.GetFocusList(ctx, userId)
}
