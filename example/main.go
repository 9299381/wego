package example

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/example/src/provider"
	"github.com/9299381/wego/example/src/router"
	"github.com/9299381/wego/providers"
)

func main()  {


	wego.Provider(&providers.BootStrap{})

	wego.Provider(&provider.ExamProvider{})


	wego.Router("grpc",&router.GrpcRouter{})
	wego.Router("http",&router.HttpRouter{})
	wego.Router("queue",&router.QueueRouter{})
	wego.Router("command",&router.CommandRouter{})
	wego.Router("websocket",&router.WebSocketRouter{})
	wego.Router("timer",&router.TimerRouter{})
	wego.Router("cron",&router.CronRouter{})


	wego.Start()



}

