package server

import (
	"context"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/domain"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/kitex_gen/api"
)

// SpaceImpl implements the last domain interface defined in the IDL.
type SpaceImpl struct {
	dc domain.DynamicUseCase
}

func NewSpaceImpl(dc domain.DynamicUseCase) api.Space {
	return &SpaceImpl{dc: dc}
}

// GetFocusDynamic implements the SpaceImpl interface.
func (s *SpaceImpl) GetFocusDynamic(ctx context.Context, req *api.SpaceGetFocusDynamicReq) (resp *api.SpaceGetFocusDynamicReply, err error) {
	// TODO: Your code here...
	return
}

// PushDynamic implements the SpaceImpl interface.
func (s *SpaceImpl) PushDynamic(ctx context.Context, req *api.SpacePushDynamicReq) (resp *api.SpacePushDynamicReply, err error) {
	// TODO: Your code here...
	return
}
