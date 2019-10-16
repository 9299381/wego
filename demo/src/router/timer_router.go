package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type TimerRouter struct {
	*servers.TimerCommServer
}

func (s *TimerRouter) Boot() {
	s.TimerCommServer = servers.NewTimerCommServer()
}

func (s *TimerRouter) Register() {

	params := make(map[string]interface{})
	params["timer"] = "test"
	s.Route("one", 2, wego.Handler("one"), params)
	s.Route("two", 5, wego.Handler("two"), params)

}
