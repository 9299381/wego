package commands

import (
	"context"
	"errors"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/commons"
)

type Server struct {
	handlers map[string]*commons.CommHandler
	Logger   contracts.ILogger
}

func NewServer() *Server {
	//初始化,logger,redis池
	s := &Server{
		handlers: make(map[string]*commons.CommHandler),
	}
	return s
}

func (s *Server) Register(name string, handler *commons.CommHandler) {
	s.handlers[name] = handler

}

func (s *Server) Serve() error {
	if args.Cmd != "" {
		//调用服务
		handler, isExist := s.handlers[args.Cmd]
		if isExist == false {
			return errors.New(constants.ErrRoute)
		}
		ctx := context.Background()
		response, err := handler.Handle(ctx, args.Args)
		if err != nil {
			return err
		}
		s.Logger.Info(response)
	}
	return nil
}
func (s *Server) Close() {

}
