// Code generated by Kitex v0.11.3. DO NOT EDIT.
package consumer

import (
	server "github.com/cloudwego/kitex/server"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/kitex_gen/api"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.Consumer, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler api.Consumer, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
