package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

//websocket 服务器尽量采用 emqx mqtt broker
type WebSocketRouter struct {
	*servers.WebSocketCommServer
}

func (s *WebSocketRouter) Boot() {
	s.WebSocketCommServer = servers.NewWebSocketCommServer()
}

func (s *WebSocketRouter) Register() {
	s.Route("Two", wego.Handler("two"))
}
