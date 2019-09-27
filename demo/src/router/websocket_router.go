package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type WebSocketRouter struct {
	*servers.WebSocketCommServer
}

func (it *WebSocketRouter) Boot() {
	it.WebSocketCommServer = servers.NewWebSocketCommServer()
}

func (it *WebSocketRouter) Register() {
	it.Route("Two", wego.Handler("two"))
}
