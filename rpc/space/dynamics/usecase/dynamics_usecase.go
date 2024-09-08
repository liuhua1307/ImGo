package usecase

import (
	"context"
	"github.com/liuhua1307/ImGo/rpc/space/domain"
)

type DynamicsUseCase struct {
	dynamicsRepo domain.DynamicRepository
}

func (d DynamicsUseCase) GetFocusDynamic(ctx context.Context, targetId, limit int64) (domain.DynamicInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (d DynamicsUseCase) PushDynamic(ctx context.Context, info domain.DynamicInfo) error {
	//TODO implement me
	panic("implement me")
}

func NewDynamicsUseCase(dynamicsRepo domain.DynamicRepository) domain.DynamicUseCase {
	return &DynamicsUseCase{dynamicsRepo: dynamicsRepo}
}
