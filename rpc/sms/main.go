package main

import (
	api "github.com/liuhua1307/ImGo/rpc/sms/kitex_gen/api/sms"
	"log"
)

func main() {
	svr := api.NewServer(new(SmsImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
