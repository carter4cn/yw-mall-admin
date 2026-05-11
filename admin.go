package main

import (
	"flag"
	"fmt"

	"mall-admin-api/internal/config"
	"mall-admin-api/internal/handler"
	"mall-admin-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting mall-admin-api at %s:%d ...\n", c.Host, c.Port)
	server.Start()
}
