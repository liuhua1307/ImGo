package main

import (
	"github.com/liuhua1307/ImGo/rpc/space/configs"
	"github.com/liuhua1307/ImGo/rpc/space/consts"
	"log"
)

func main() {
	configs.Init()
	configs.LogInit(consts.DefaultLogFilePath)
	svr := wireApp()
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
