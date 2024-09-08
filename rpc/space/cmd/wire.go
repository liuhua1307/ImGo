//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	kitexServer "github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/space/configs"
	cacheRepo "github.com/liuhua1307/ImGo/rpc/space/dynamics/cache"
	"github.com/liuhua1307/ImGo/rpc/space/dynamics/repository"
	"github.com/liuhua1307/ImGo/rpc/space/dynamics/usecase"
	"github.com/liuhua1307/ImGo/rpc/space/pkg/cache"
	"github.com/liuhua1307/ImGo/rpc/space/server"
)

func wireApp() kitexServer.Server {
	panic(wire.Build(server.ProviderSet, cache.ProviderSet, repository.ProviderSet, usecase.ProviderSet, configs.ProviderSet, cacheRepo.ProviderSet))
}
