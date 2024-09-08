package main

import (
	"context"
	api "github.com/liuhua1307/ImGo/rpc/user/kitex_gen/api"
)

// UserImpl implements the last domain interface defined in the IDL.
type UserImpl struct{}

// Register implements the UserImpl interface.
func (s *UserImpl) Register(ctx context.Context, req *api.UserRegisterReq) (resp *api.UserRegisterReply, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserImpl interface.
func (s *UserImpl) Login(ctx context.Context, req *api.UserLoginReq) (resp *api.UserLoginReply, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserImpl interface.
func (s *UserImpl) GetUserInfo(ctx context.Context, req *api.UserGetUserInfoReq) (resp *api.UserGetUserInfoReply, err error) {
	// TODO: Your code here...
	return
}
