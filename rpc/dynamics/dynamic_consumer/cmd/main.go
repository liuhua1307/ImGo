package main

import (
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/configs"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_consumer/consts"
)

func main() {
	configs.LogInit(consts.DefaultLogFilePath)
	app := wireApp()
	go func() {
		err := configs.NewConsumer().Start()
		if err != nil {
			panic(err)
		}
	}()
	if err := app.Run(); err != nil {
		panic(err)
	}

}
