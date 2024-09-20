package main

import (
	"alger/common/context"
	"alger/common/global"
	"alger/common/utils/env"
	"alger/config"
	"alger/middleware/recover"
	"alger/router"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"os"
)

func main() {
	//加载配置文件
	config := loadConfig()
	//初始化必要组件
	global.InitComponents(config)
	//启动组件
	enableComponents(config)
	//优雅启动多个服务
	group := service.NewServiceGroup()
	defer group.Stop()
	group.Add(newHttpService(config))
	group.Start()
}

func enableComponents(config *config.Config) {
	context.NewApiServiceContext(*config)

}

const (
	REMOTEPROVIDER string = "consul"
	CONSULURL      string = "CONSUL_URL"
	PATH           string = "alger" //consul-> key
)

func loadConfig() *config.Config {
	var c config.Config
	consulUrl := os.Getenv(CONSULURL)
	if len(consulUrl) == 0 {
		fmt.Println("use local config")
		readLocalConfig(&c)
	} else {
		fmt.Println("use config center, ", consulUrl)
		readRemoteConfig(&c)
	}
	return &c
}

func readLocalConfig(c *config.Config) {
	v := viper.New()
	v.AddConfigPath(env.CheckPath())
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	err := v.Unmarshal(c)
	if err != nil {
		panic(err)
	}
}

func readRemoteConfig(c *config.Config) {
	err := viper.AddRemoteProvider(REMOTEPROVIDER, os.Getenv(CONSULURL), PATH)
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("yaml")

	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(c)
	if err != nil {
		panic(err)
	}
}

type http struct {
	config config.Config
	server *rest.Server
}

func newHttpService(config *config.Config) *http {
	return &http{config: *config}
}

func (h *http) Start() {
	ctx := context.NewApiServiceContext(h.config)
	server := rest.MustNewServer(h.config.Api.RestConf)
	//全局recover
	server.Use(recover.NewHttpRecover().Handle)
	h.server = server
	router.RegisterHandlers(server, ctx)
	global.Logger.Info("starting api server at %s:%d...\n", h.config.Api.Host, h.config.Api.Port)

	server.Start()
}

func (h *http) Stop() {
	global.Logger.Info("stop api server")
	h.server.Stop()
}
