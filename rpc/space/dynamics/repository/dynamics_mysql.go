package repository

import (
	"context"
	"github.com/liuhua1307/ImGo/rpc/space/domain"
	"gorm.io/gorm"
)

type DynamicsMYSQLRepository struct {
	mysqlDB *gorm.DB
}

func NewDynamicsMYSQLRepository(mysqlDB *gorm.DB) *DynamicsMYSQLRepository {
	return &DynamicsMYSQLRepository{mysqlDB: mysqlDB}
}

func (d DynamicsMYSQLRepository) GetFocusDynamic(ctx context.Context, targetId, limit, userId int64) ([]domain.Dynamics, error) {
	//TODO implement me
	panic("implement me")
}

func (d DynamicsMYSQLRepository) PushDynamicAndPush(ctx context.Context, info domain.DynamicInfo) error {
	//TODO implement me
	panic("implement me")
}
