package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type SubscribeRouter struct {
	*servers.MqttSubscribeCommCommServer
}

func (it *SubscribeRouter) Boot() {
	it.MqttSubscribeCommCommServer = servers.NewMqttSubscribeCommCommServer()
}

func (it *SubscribeRouter) Register() {
	//topic -> handler
	it.Route("sub_test", wego.Handler("two"))
	it.Route("sub_test2", wego.Handler("two"))
	it.Route("sub_test3", wego.Handler("sleep"))

}
