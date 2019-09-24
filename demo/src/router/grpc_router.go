package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type GrpcRouter struct {
	*servers.GrpcCommServer
}

func (it *GrpcRouter) Boot() {
	it.GrpcCommServer = servers.NewGrpcCommServer()
}

//这里注册路由
func (it *GrpcRouter) Register() {

	it.Route("demo.two", wego.Handler("two"))
	it.Route("demo.one", wego.Handler("one"))

}
