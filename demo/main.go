package main

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/demo/src/provider"
	"github.com/9299381/wego/demo/src/router"
	"github.com/9299381/wego/providers"
	"strings"
)

func main() {
	wego.Provider(&providers.BootStrap{})
	wego.Provider(&provider.ExamProvider{})

	servers := strings.Split(args.Server, ",")
	for _, s := range servers {
		if s == "http" {
			wego.Router("http", &router.HttpRouter{})
		}
		if s == "grpc" {
			wego.Router("grpc", &router.GrpcRouter{})
		}
		if s == "queue" {
			wego.Router("queue", &router.QueueRouter{})
		}
		if s == "websocket" {
			wego.Router("websocket", &router.WebSocketRouter{})
		}
		if s == "command" {
			wego.Router("command", &router.CommandRouter{})
		}
		if s == "timer" {
			wego.Router("timer", &router.TimerRouter{})
		}
		if s == "cron" {
			wego.Router("cron", &router.CronRouter{})
		}
	}

	wego.Start()

}
