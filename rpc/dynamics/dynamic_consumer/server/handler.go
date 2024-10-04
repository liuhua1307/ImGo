package server

import (
	"context"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/domain"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/kitex_gen/api"
)

// ConsumerImpl implements the last service interface defined in the IDL.
type ConsumerImpl struct {
	dc domain.DynamicUseCase
}

func NewConsumerImpl(dc domain.DynamicUseCase) api.Consumer {
	return &ConsumerImpl{dc: dc}
}

// PushDynamic implements the ConsumerImpl interface.
func (s *ConsumerImpl) PushDynamic(ctx context.Context, req *api.ConsumerPushDynamicReq) (resp *api.ConsumerPushDynamicReply, err error) {
	// TODO: Your code here...
	return
}
