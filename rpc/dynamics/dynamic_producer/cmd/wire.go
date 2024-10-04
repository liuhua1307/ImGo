//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	kitexServer "github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/configs"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/pkg/cache"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/pkg/mq"
	cacheRepo "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/producer/cache"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/producer/repository"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/producer/usecase"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/server"
)

func wireApp() kitexServer.Server {
	panic(wire.Build(server.ProviderSet, cache.ProviderSet, repository.ProviderSet, usecase.ProviderSet, configs.ProviderSet, cacheRepo.ProviderSet, mq.ProviderSet))
}
