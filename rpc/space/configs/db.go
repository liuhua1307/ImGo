package configs

import (
	"context"
	"fmt"
	"github.com/liuhua1307/ImGo/rpc/space/consts"
	pkgLog "github.com/liuhua1307/ImGo/rpc/space/pkg/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var (
	db     *gorm.DB
	onceDb sync.Once
)

func DB(ctx context.Context) *gorm.DB {
	onceDb.Do(func() {
		var err error
		if db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: fmt.Sprintf("incu_user:%v@tcp(localhost:3307)/incu_om?charset=utf8mb4&parseTime=true&loc=Local", viper.GetString("INCU_PASSWORD")),
		}), &gorm.Config{
			Logger:      getGormLogger(),
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}); err != nil {
			panic(err)
		}

	})
	return db.WithContext(ctx)
}

func getGormLogger() logger.Interface {
	ignoreRecordNotFound := false
	logLevel := logger.Info
	if !viper.GetBool("debug") {
		ignoreRecordNotFound = true
		logLevel = logger.Error
	}
	logFile, err := os.OpenFile(consts.DefaultSQLLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	multiOutput := io.MultiWriter(os.Stdout, logFile)
	return logger.New(
		log.New(multiOutput, "[DB] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: ignoreRecordNotFound,
			Colorful:                  false,
		},
	)
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		pkgLog.Log().Error("get sql db error", pkgLog.Field{Key: "err", Value: err})
	}
	if err = sqlDB.Close(); err != nil {
		pkgLog.Log().Error("sql db close error", pkgLog.Field{Key: "err", Value: err})
	}
}
