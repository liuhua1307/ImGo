package configs

import (
	"context"
	"github.com/google/wire"
	"github.com/liuhua1307/ImGo/rpc/space/consts"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
)

var ProviderSet = wire.NewSet(DB, NewRDB, NewContext)

type Option func(*conf)

func NewContext() context.Context {
	return context.Background()
}

type conf struct {
	configFileType string
	configFilename string
}

// c: immutable default config
var c = conf{
	configFileType: consts.DefaultConfigFileType,
	configFilename: consts.DefaultConfigFileName,
}

func apply(opts ...Option) *conf {
	newConf := c
	for _, opt := range opts {
		opt(&newConf)
	}
	return &newConf
}

func WithFileType(fileType string) Option {
	return func(c *conf) {
		c.configFileType = fileType
	}
}

func WithFileName(filename string) Option {
	return func(c *conf) {
		c.configFilename = filename
	}
}

func Init(opts ...Option) {
	cur := apply(opts...)
	viper.AutomaticEnv()

	pflag.String("config", "config.yaml", "configuration file path")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic("failed to parse command line flags")
	}

	configFile := viper.GetString("config")
	if len(configFile) == 0 {
		configFile = "config.yaml"
	}

	if _, err := os.Stat(configFile); err != nil {
		log.Fatalln("config.yaml not found, please add `--config`")
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType(cur.configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
