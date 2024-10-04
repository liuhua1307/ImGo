package configs

import log2 "github.com/liuhua1307/ImGo/rpc/dynamics/dynamic_producer/pkg/log"

func LogInit(path string) {
	log2.RegisterLog(log2.NewSlogLogger(path))
}
