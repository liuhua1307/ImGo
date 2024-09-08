package main

import (
	api "github.com/liuhua1307/ImGo/rpc/user/kitex_gen/api/user"
	"log"
)

func main() {
	svr := api.NewServer(new(UserImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
