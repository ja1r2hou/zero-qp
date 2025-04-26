package main

import (
	"flag"
	"fmt"
	"zero-qp/wsService/api/internal/config"
	"zero-qp/wsService/api/internal/handler"
	"zero-qp/wsService/api/internal/net"
	"zero-qp/wsService/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/wsApi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	go func() {
		net.InitWs(ctx)
	}()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
