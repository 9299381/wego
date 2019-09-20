package servers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/transports"
	"github.com/9299381/wego/servers/websockets"
	"github.com/go-kit/kit/endpoint"
)

type WebSocketCommServer struct {
	*websockets.Server
}

func NewWebSocketCommServer() *WebSocketCommServer {

	ss:= &WebSocketCommServer{
		Server: websockets.NewServer(),
	}
	ss.Server.Logger = wego.App.Logger
	return ss
}

func (it *WebSocketCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewWebSocket(endpoint),
	}
	it.Register(name, handler)
}

func (it *WebSocketCommServer)Load()  {
	//注册通用路由
}


func (it *WebSocketCommServer) Start() error {
	return it.Serve()
}