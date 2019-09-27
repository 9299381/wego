package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type QueueRouter struct {
	*servers.QueueCommServer
}

func (it *QueueRouter) Boot() {
	it.QueueCommServer = servers.NewQueueCommServer()
}

func (it *QueueRouter) Register() {
	it.Route("queue_test", wego.Handler("two"))
}
