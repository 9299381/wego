package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type CommandRouter struct {
	*servers.CommandCommServer
}

func (it *CommandRouter) Boot()  {
	it.CommandCommServer = servers.NewCommandCommServer()
}

//todo 写个队列server 编解码,路由等

func (it *CommandRouter) Register()  {
	it.Route("cmd_test", wego.Handler("two"))
}