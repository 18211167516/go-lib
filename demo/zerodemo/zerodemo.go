package main

import (
	"flag"
	"fmt"

	"github.com/18211167516/go-lib/demo/zerodemo/internal/config"
	"github.com/18211167516/go-lib/demo/zerodemo/internal/handler"
	"github.com/18211167516/go-lib/demo/zerodemo/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/zerodemo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
