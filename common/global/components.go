package global

import (
	"alger/common/gorm"
	"alger/common/i18n"

	"alger/common/log"
	"alger/common/utils/env"
	"alger/config"
	"github.com/zeromicro/go-zero/core/logx"
	redisPkg "github.com/zeromicro/go-zero/core/stores/redis"

	gormPkg "gorm.io/gorm"
)

var Logger log.Log
var Db *gormPkg.DB
var Sql map[string]string
var Redis *redisPkg.Redis
var Bundle *i18n.Bundle

var Configuration *config.Config

func InitComponents(c *config.Config) {

	Configuration = c
	//关闭原生日志
	logx.DisableStat()

	//初始化log
	Logger = log.Initialize(c.Log)

	//启动Gorm支持
	var err error
	Db, err = gorm.Initialize(c.Mysql)
	if err != nil {
		panic(err)
	}
	Bundle, err = i18n.Initialize(env.CheckPath() + "/i18n/")
	if err != nil {
		panic(err)
	}
}

const WORKIDBITNUM = 1024
