package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type SubscribeRouter struct {
	*servers.MqttSubscribeCommCommServer
}

func (s *SubscribeRouter) Boot() {
	s.MqttSubscribeCommCommServer = servers.NewMqttSubscribeCommCommServer()
}

func (s *SubscribeRouter) Register() {
	//topic -> handler
	s.Route("sub_test", wego.Handler("two"))
	s.Route("sub_test2", wego.Handler("two"))
	s.Route("sub_test3", wego.Handler("sleep"))

}
