package main

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/demo/src/provider"
	"github.com/9299381/wego/demo/src/router"
	"github.com/9299381/wego/providers"
	"github.com/9299381/wego/servers"
)

func main() {

	//args.Registy = "127.0.0.1:8500"
	args.Server = "http"

	//args.Server = "http,event,subscribe"
	//args.Server = "subscribe"
	//args.Server = "queue"
	args.Name = "demo"
	args.Mode = "dev"
	//服务注册
	wego.Provider(&providers.ConsulRegistyProvider{})
	// api 接口
	wego.Provider(&provider.DemoProvider{})
	// http服务器路由
	wego.Router("http", &router.HttpRouter{})
	// grpc_api 接口服务路由
	wego.Router("grpc", &router.GrpcRouter{})
	wego.Router("queue", &router.QueueRouter{})
	wego.Router("command", &router.CommandRouter{})
	wego.Router("timer", &router.TimerRouter{})
	wego.Router("cron", &router.CronRouter{})
	wego.Router("websocket", &router.WebSocketRouter{})
	wego.Router("subscribe", &router.SubscribeRouter{})
	//内置加载事件服务,无需路由,直接调用 filter handler
	wego.Router("event", servers.NewEventCommServer())

	wego.Start()

}
