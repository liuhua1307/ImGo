package main

import (
	api "github.com/liuhua1307/ImGo/rpc/chat/kitex_gen/api/chat"
	"log"
)

func main() {
	svr := api.NewServer(new(ChatImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
