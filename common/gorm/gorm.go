package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DataSource string
	LogMode    bool
}

func Initialize(config Config) (db *gorm.DB, err error) {
	//启动Gorm支持
	var logMode logger.LogLevel
	if config.LogMode {
		logMode = logger.Info
	} else {
		logMode = logger.Silent
	}
	return gorm.Open(mysql.Open(config.DataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
}
