package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type CronRouter struct {
	*servers.CronCommServer
}

func (s *CronRouter) Boot() {
	s.CronCommServer = servers.NewCronCommServer()
}

func (s *CronRouter) Register() {

	s.Route("*/5 * * * * *", wego.Handler("one"))
	s.Route("*/2 * * * * *", wego.Handler("two"))
}
