package servers

import (
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/commands"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

type CommandCommServer struct {
	*commands.Server
}

func NewCommandCommServer() *CommandCommServer {
	ss := &CommandCommServer{
		Server: commands.NewServer(),
	}
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *CommandCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewCommand(endpoint),
	}
	s.Register(name, handler)
}

func (s *CommandCommServer) Load() {

	//注册通用路由
}

func (s *CommandCommServer) Start() error {
	return s.Serve()

}
func (s *CommandCommServer) Close() {
	s.Server.Close()
}
