package usecase

import "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/domain"

type DynamicsUseCase struct {
	dynamicsRepo domain.DynamicRepository
}

func NewDynamicsUseCase(dynamicsRepo domain.DynamicRepository) domain.DynamicUseCase {
	return &DynamicsUseCase{dynamicsRepo: dynamicsRepo}
}

func (d DynamicsUseCase) AsyncPushDynamic(info domain.Dynamic) error {
	return d.dynamicsRepo.PushDynamic(info)
}
