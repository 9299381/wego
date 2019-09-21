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

//todo 写个队列server 编解码,路由等

func (it *WebSocketRouter) Register() {
	it.Route("websocket_test", wego.Handler("two"))
}
