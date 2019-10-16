package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type CommandRouter struct {
	*servers.CommandCommServer
}

func (s *CommandRouter) Boot() {
	s.CommandCommServer = servers.NewCommandCommServer()
}

//todo 写个队列server 编解码,路由等

func (s *CommandRouter) Register() {
	s.Route("cmd_test", wego.Handler("two"))
}
