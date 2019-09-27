package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type TimerRouter struct {
	*servers.TimerCommServer
}

func (it *TimerRouter) Boot() {
	it.TimerCommServer = servers.NewTimerCommServer()
}

func (it *TimerRouter) Register() {

	params := make(map[string]interface{})
	params["timer"] = "test"
	it.Route("one", 2, wego.Handler("one"), params)
	it.Route("two", 5, wego.Handler("two"), params)

}
