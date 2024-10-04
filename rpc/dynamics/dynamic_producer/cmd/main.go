package main

import (
	configs2 "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/configs"
	"github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/consts"
)

func main() {
	configs2.Init()
	configs2.LogInit(consts.DefaultLogFilePath)
	svr := wireApp()
	go func() {
		err := configs2.NewProducer().Start()
		if err != nil {
			panic(err)
		}
	}()
	err := svr.Run()
	if err != nil {
		panic(err)
	}
}
