package servers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/events"
	"time"
)

type EventCommServer struct {
	*events.Server
}

func NewEventCommServer(freq int) *EventCommServer {
	ss := &EventCommServer{
		Server: events.NewServer(),
	}
	ss.Ticker = time.NewTicker(time.Duration(freq) * time.Second)
	return ss
}
func (it *EventCommServer) Boot() {
	event := contracts.NewEvent()
	wego.App.Event = event
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
