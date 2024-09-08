package main

import (
	"context"
	api "github.com/liuhua1307/ImGo/rpc/chat/kitex_gen/api"
)

// ChatImpl implements the last domain interface defined in the IDL.
type ChatImpl struct{}

// Send implements the ChatImpl interface.
func (s *ChatImpl) Send(ctx context.Context, req *api.ChatSendReq) (resp *api.ChatSendReply, err error) {
	// TODO: Your code here...
	return
}

// GetMessageList implements the ChatImpl interface.
func (s *ChatImpl) GetMessageList(ctx context.Context, req *api.ChatGetMessageListReq) (resp *api.ChatGetMessageListReply, err error) {
	// TODO: Your code here...
	return
}
