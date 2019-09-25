package main

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/demo/src/provider"
	"github.com/9299381/wego/demo/src/router"
	"github.com/9299381/wego/providers"
)

func main() {

	args.Registy = "127.0.0.1:8500"
	args.Server = "http,grpc,event"
	args.Name = "consul_demo"

	wego.Provider(&providers.BootStrap{})
	wego.Provider(&provider.DemoProvider{})

	wego.Router("http", &router.HttpRouter{})
	wego.Router("grpc", &router.GrpcRouter{})
	wego.Router("queue", &router.QueueRouter{})
	wego.Router("command", &router.CommandRouter{})
	wego.Router("timer", &router.TimerRouter{})
	wego.Router("cron", &router.CronRouter{})
	wego.Router("websocket", &router.WebSocketRouter{})

	wego.Start()

}
