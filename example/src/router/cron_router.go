package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type CronRouter struct {
	*servers.CronCommServer
}

func (it *CronRouter) Boot()  {
	it.CronCommServer = servers.NewCronCommServer()
}

//todo 写个队列server 编解码,路由等

func (it *CronRouter) Register()  {

	it.Route("*/5 * * * * *", wego.Handler("one"))
	it.Route("*/2 * * * * *", wego.Handler("two"))
}