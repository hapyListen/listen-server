package main

import (
	"flag"
	"fmt"

	"listen-server/app/user/cmd/api/user/internal/config"
	"listen-server/app/user/cmd/api/user/internal/handler"
	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/common/handle"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// Set the error handling function
	httpx.SetErrorHandler(handle.ErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
