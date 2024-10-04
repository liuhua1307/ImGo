package repository

import (
	"context"
	domain2 "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/domain"
	"gorm.io/gorm"
)

type DynamicsMYSQLRepository struct {
	mysqlDB *gorm.DB
}

func NewDynamicsMYSQLRepository(mysqlDB *gorm.DB) *DynamicsMYSQLRepository {
	return &DynamicsMYSQLRepository{mysqlDB: mysqlDB}
}

func (d DynamicsMYSQLRepository) GetFocusDynamic(ctx context.Context, targetId, limit, userId int64) ([]domain2.Dynamics, error) {
	var dynamics []domain2.Dynamics
	subQuery := d.mysqlDB.Model(&domain2.Focus{}).Select("target_id").Where("user_id = ?", userId)
	err := d.mysqlDB.Model(&domain2.Dynamics{}).Select("author_id", "title", "info", "image_urls", "created_at", "deleted_at").Where("author_id = (?)", subQuery).Limit(int(limit)).Find(&dynamics).Error
	return dynamics, err
}

func (d DynamicsMYSQLRepository) PushDynamicAndPush(ctx context.Context, info domain2.DynamicInfo) error {
	//TODO implement me
	panic("implement me")
}

func (d DynamicsMYSQLRepository) GetFocusList(ctx context.Context, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}
