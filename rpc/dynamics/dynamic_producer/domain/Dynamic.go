package domain

import (
	"context"
	"time"
)

type DynamicUseCase interface {
	GetFocusDynamic(ctx context.Context, targetId, limit int64) (DynamicInfo, error)
	PushDynamic(ctx context.Context, info DynamicInfo) error
}
type DynamicRepository interface {
	GetFocusDynamic(ctx context.Context, targetId, limit, userId int64) ([]Dynamics, error)
	PushDynamicAndPush(ctx context.Context, info DynamicInfo) error
	GetFocusList(ctx context.Context, userId int64) ([]int64, error)
}

type Dynamics struct {
	ID        uint64
	AuthorId  uint64
	Title     string
	Info      string
	ImageUrls string
	CreatedAt time.Time
	DeletedAt time.Time
}

type DynamicInfo struct {
	ID          uint64
	AuthorId    uint64
	Title       string
	Info        string
	ImageUrl    []string
	CommentList []uint64
}
