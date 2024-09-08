package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/liuhua1307/ImGo/rpc/space/domain"
	"github.com/liuhua1307/ImGo/rpc/space/dynamics/repository"
	"github.com/liuhua1307/ImGo/rpc/space/pkg/cache"
	"strconv"
	"time"
)

const UserTimeLineProfix = "UserTimeLine-"

type DynamicsRedisRepository struct {
	dynamicsDBRepository domain.DynamicRepository
	pool                 cache.Cache
}

func NewDynamicsRepository(pool cache.Cache, dbRepo *repository.DynamicsMYSQLRepository) *DynamicsRedisRepository {
	return &DynamicsRedisRepository{pool: pool, dynamicsDBRepository: dbRepo}
}

func (h DynamicsRedisRepository) GetFocusDynamic(ctx context.Context, targetId, limit, userId int64) ([]domain.Dynamics, error) {
	card, err := h.pool.ZCard(UserTimeLineProfix + strconv.FormatInt(userId, 10))
	if err == redis.Nil || (err == nil && card <= 1) {
		return h.getFocusDynamicWithPull(ctx, tar																																																																																																																																																																																																																																																																																																																																																																	getId, limit, userId)
	}
	if err != nil {
		return nil, err
	}
}

func (h DynamicsRedisRepository) GetFocusDynamicWithoutPull(ctx context.Context, userId int64) ([]domain.Dynamics, error) {
	//TODO implement me
	panic("implement me")
}

func (h DynamicsRedisRepository) PushDynamicAndPush(ctx context.Context, info domain.DynamicInfo) error {
	//TODO implement me
	panic("implement me")
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
	err = h.pool.ZSet(UserTimeLineProfix+strconv.FormatInt(userId, 10), redisZ, time.Hour*72)
	if err != nil {
		return nil, err
	}

	return dynamics, nil
}
