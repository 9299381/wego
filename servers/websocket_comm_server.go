package servers

import (
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/transports"
	"github.com/9299381/wego/servers/websockets"
	"github.com/go-kit/kit/endpoint"
)

//websocket 服务器尽量采用 emqx mqtt broker
type WebSocketCommServer struct {
	*websockets.Server
}

func NewWebSocketCommServer() *WebSocketCommServer {

	ss := &WebSocketCommServer{
		Server: websockets.NewServer(),
	}
	ss.Logger = loggers.GetLog()
	return ss
}

func (it *WebSocketCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewWebSocket(endpoint),
	}
	it.Register(name, handler)
}

func (it *WebSocketCommServer) Load() {
	//注册通用路由
}

func (it *WebSocketCommServer) Start() error {
	return it.Serve()
}
func (it *WebSocketCommServer) Close() {
	it.Server.Close()
}
