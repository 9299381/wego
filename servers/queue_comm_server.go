package servers

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/queues"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

/**
redis queue
*/
type QueueCommServer struct {
	*queues.Server
}

func NewQueueCommServer() *QueueCommServer {
	config := configs.LoadQueueConfig()
	ss := &QueueCommServer{
		Server: queues.NewServer(&queues.Options{
			Prefix:      config.Prefix,
			Listen:      config.Listen,
			Interval:    config.Interval,
			UseNumber:   true,
			Concurrency: config.Concurrency,
		}),
	}
	ss.RedisPool = clients.RedisPool()
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *QueueCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewQueue(endpoint),
	}
	s.Register(name, handler)
}

func (s *QueueCommServer) Load() {

	//注册通用路由
}

func (s *QueueCommServer) Start() error {
	return s.Serve()

}
func (s *QueueCommServer) Close() {
	s.Server.Close()
}
