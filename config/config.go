package config

import (
	"alger/common/gorm"
	"alger/common/log"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Api ApiConfig

	Log   log.Config
	Mysql gorm.Config
}

// Sms
// @Description: 短信配置
//
//	type Sms struct {
//		sms.Config    `mapstructure:",squash"`
//		AlertsCodeTmp string //告警
//	}
type ApiConfig struct {
	rest.RestConf `mapstructure:",squash"`
	//UserRpc       zrpc.RpcClientConf
	//AssetRpc      zrpc.RpcClientConf
	//GatewayRpc    zrpc.RpcClientConf
	//MetadataRpc   zrpc.RpcClientConf
}

type RpcConfig struct {
	zrpc.RpcServerConf `mapstructure:",squash"`
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
