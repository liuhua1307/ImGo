package main

import (
	"context"
	api "github.com/liuhua1307/ImGo/rpc/sms/kitex_gen/api"
)

// SmsImpl implements the last domain interface defined in the IDL.
type SmsImpl struct{}

// GetSMSCode implements the SmsImpl interface.
func (s *SmsImpl) GetSMSCode(ctx context.Context, req *api.SmsGetCodeReq) (resp *api.SmsGetCodeReply, err error) {
	// TODO: Your code here...
	return
}

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
