package configs

import "github.com/liuhua1307/ImGo/rpc/space/pkg/log"

func LogInit(path string) {
	log.RegisterLog(log.NewSlogLogger(path))
}
