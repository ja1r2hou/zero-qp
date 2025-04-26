package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	LogConf logx.LogConf
	Jwt     struct {
		Secret string `mapstructure:"secret"`
		Exp    int64  `mapstructure:"exp"`
	}
}
