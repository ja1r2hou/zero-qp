package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-qp/userService/rpc/user"
	"zero-qp/userService/rpc/userrpc"
	"zero-qp/wsService/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	User   user.UserRPCClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userrpc.NewUserRPC(zrpc.MustNewClient(c.UserRpc)),
	}
}
