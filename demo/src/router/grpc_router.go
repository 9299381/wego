package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type GrpcRouter struct {
	*servers.GrpcCommServer
}

func (s *GrpcRouter) Boot() {
	s.GrpcCommServer = servers.NewGrpcCommServer()
}

//这里注册路由
func (s *GrpcRouter) Register() {

	s.Route("demo.two", wego.Handler("two"))
	s.Route("demo.one", wego.Handler("one"))
	s.Route("demo.post", wego.Handler("post"))

}
