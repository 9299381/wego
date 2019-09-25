package servers

import (
	"github.com/9299381/wego/filters"
	"github.com/9299381/wego/servers/gateway"
	"github.com/go-kit/kit/endpoint"
)

type GateWayCommServer struct {
	*gateway.Server
}

func NewGateWayCommServer() *GateWayCommServer {
	return &GateWayCommServer{
		Server: gateway.NewServer(),
	}
}

func (it *GateWayCommServer) Route(method, path string, endpoint endpoint.Endpoint) {
	//如果有本地注册的路由,则跑本地,2种情况组合endpoint filter
	//1 跑本地服务
	//2 只跑本地endpoint filter
	it.Register(method, path, endpoint)
}

func (it *GateWayCommServer) Load() {
	//注册通用路由,consul 心跳检测
	it.Route("GET", "/health", (&filters.HealthEndpoint{}).Make())

}

func (it *GateWayCommServer) Start() error {
	return it.Serve()
}
