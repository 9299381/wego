package servers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/events"
	"time"
)

type EventCommServer struct {
	*events.Server
}

func NewEventCommServer() *EventCommServer {
	freq := 3 //从config中读取
	ss := &EventCommServer{
		Server: events.NewServer(),
	}
	ss.Logger = loggers.Log
	ss.Ticker = time.NewTicker(time.Duration(freq) * time.Second)
	events.Handlers = wego.App.Handlers
	return ss
}
func (it *EventCommServer) Boot() {

}

func (it *EventCommServer) Load() {

	//注册通用路由
}
func (it *EventCommServer) Register() {
}

func (it *EventCommServer) Route() {

}

func (it *EventCommServer) Start() error {
	return it.Serve()
}
