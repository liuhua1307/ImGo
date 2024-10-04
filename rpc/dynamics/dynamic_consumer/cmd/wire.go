//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/configs"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/consumer/cache"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/consumer/usecase"
	pkgCache "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/pkg/cache"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/pkg/mq"
	apiSever "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/server"
)

func wireApp() server.Server {
	panic(wire.Build(configs.ProviderSet, apiSever.ProviderSet, cache.ProviderSet, usecase.ProviderSet, pkgCache.ProviderSet, mq.ProviderSet))
}
