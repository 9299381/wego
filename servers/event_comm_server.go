package servers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/events"
	"time"
)

type EventCommServer struct {
	*events.Server
}

func NewEventCommServer() *EventCommServer {
	config := configs.LoadEventConfig()
	ss := &EventCommServer{
		Server: events.NewServer(),
	}
	ss.Logger = loggers.GetLog()
	ss.Concurrency = config.Concurrency
	ss.After = time.After(time.Duration(config.After) * time.Second)
	events.Handlers = wego.App.Handlers
	return ss
}
func (s *EventCommServer) Boot() {

}

func (s *EventCommServer) Load() {

	//注册通用路由
}
func (s *EventCommServer) Register() {
}

func (s *EventCommServer) Route() {

}

func (s *EventCommServer) Start() error {
	return s.Serve()
}
func (s *EventCommServer) Close() {
	s.Server.Close()
}
