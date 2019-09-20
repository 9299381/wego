package servers

import (
	"github.com/9299381/wego/servers/commands"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

type CommandCommServer struct {
	*commands.Server
}

func NewCommandCommServer()  *CommandCommServer{
	return &CommandCommServer{
		Server: commands.NewServer(),
	}
}

func (it *CommandCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewCommand(endpoint),
	}
	it.Register(name, handler)
}

func (it *CommandCommServer)Load()  {

	//注册通用路由
}

func (it *CommandCommServer) Start() error {
	return it.Serve()

}