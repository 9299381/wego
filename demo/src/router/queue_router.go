package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type QueueRouter struct {
	*servers.QueueCommServer
}

func (s *QueueRouter) Boot() {
	s.QueueCommServer = servers.NewQueueCommServer()
}

func (s *QueueRouter) Register() {
	s.Route("queue_test", wego.Handler("two"))
	s.Route("queue2", wego.Handler("queue2"))

}
