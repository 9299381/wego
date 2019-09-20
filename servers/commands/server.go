package commands

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/servers/commons"
)

type Server struct {
	handlers map[string]*commons.CommHandler
}


func NewServer() *Server {
	//初始化,logger,redis池
	s:=&Server{
		handlers:make(map[string]*commons.CommHandler),
	}
	return s
}

func (it *Server)Register(name string,handler *commons.CommHandler)  {
	it.handlers[name] = handler

}

func (it *Server)Serve() error {
	//解析指令,解析参数
	var cmd,args string
	flag.StringVar(&cmd, "cmd", "cmd","cli命令")
	flag.StringVar(&args, "args", "json","json参数")
	flag.Parse()

	if cmd != ""{
		msg :="CommandServer Stop!"
		return errors.New(msg)
	}

	//调用服务
	handler,isExist := it.handlers[cmd]
	if isExist ==false{
		return errors.New(constants.ErrRoute)
	}

	ctx := context.Background()
	response,err := handler.Handle(ctx,args)
	if err!=nil {
		return err
	}
	fmt.Println(response)
	return nil
}
