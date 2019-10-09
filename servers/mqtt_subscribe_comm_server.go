package servers

import (
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/mqtts"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

type MqttSubscribeCommCommServer struct {
	*mqtts.Server
}

func NewMqttSubscribeCommCommServer() *MqttSubscribeCommCommServer {
	ss := &MqttSubscribeCommCommServer{
		Server: mqtts.NewServer(),
	}
	ss.Logger = loggers.Log
	return ss
}

func (it *MqttSubscribeCommCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewMqttSubscribe(endpoint),
	}
	it.Register(name, handler)
}

func (it *MqttSubscribeCommCommServer) Load() {

	//注册通用路由
}

func (it *MqttSubscribeCommCommServer) Start() error {
	return it.Serve()

}
