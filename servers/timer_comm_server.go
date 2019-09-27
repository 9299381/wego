package servers

import (
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/timers"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

/**
定时器
*/
type TimerCommServer struct {
	*timers.Server
}

func NewTimerCommServer() *TimerCommServer {
	ss := &TimerCommServer{
		Server: timers.NewServer(),
	}
	ss.Logger = loggers.Log
	return ss
}

func (it *TimerCommServer) Load() {

	//注册通用路由
}

func (it *TimerCommServer) Route(name string, freq int, endpoint endpoint.Endpoint, params map[string]interface{}) {

	handler := &commons.CommHandler{
		Handler: transports.NewTimer(endpoint),
	}
	it.Register(name, freq, handler, params)
}

func (it *TimerCommServer) Start() error {
	return it.Serve()
}
