package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	RedisCof struct {
		Addr         string
		PoolSize     int
		MinIdleConns int
		Password     string
	}
	Mongo struct {
		Url         string
		MaxPoolSize int
		MinPoolSize int
		Password    string
		Db          string
		UserName    string
	}
}
