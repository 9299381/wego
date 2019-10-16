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
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *MqttSubscribeCommCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewMqttSubscribe(endpoint),
	}
	s.Register(name, handler)
}

func (s *MqttSubscribeCommCommServer) Load() {

	//注册通用路由
}

func (s *MqttSubscribeCommCommServer) Start() error {
	return s.Serve()

}

func (s *MqttSubscribeCommCommServer) Close() {
	s.Server.Close()
}
