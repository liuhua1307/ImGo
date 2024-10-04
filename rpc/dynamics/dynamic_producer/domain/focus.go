package domain

import "time"

type Focus struct {
	ID        uint64
	UserId    uint64
	TargetId  uint64
	CreatedAt time.Time
	DeletedAt time.Time
}
