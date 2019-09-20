package timers

import (
	"context"
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers/commons"
	"github.com/sirupsen/logrus"
	"time"
)

type Server struct {
	handlers map[string] *service
	Logger *logrus.Logger
}
type service struct {
	freq int
	handler *commons.CommHandler
	params map[string]interface{}
}

func NewServer() *Server {
	ss := &Server{
		handlers: make(map[string] *service),
	}

	return ss
}
func (it *Server) Register(name string,freq int,handler *commons.CommHandler,params map[string]interface{}){

	it.handlers[name] = &service{
		freq:freq,
		handler:handler,
		params:params,
	}
}



func (it *Server) Serve() error  {

	errChans := make(map[string]chan error)
	for name,svr := range it.handlers {
		errChans[name] = make(chan error)
		ticker := time.NewTicker(time.Duration(svr.freq) * time.Second)
		go func(name string,svr *service,t *time.Ticker,errChan chan error) {
			for {
				select {
				case <-t.C:
					id := wego.ID()
					log := it.Logger.WithFields(logrus.Fields{
						"request_id" : id,
					})
					log.Info(name + ":任务开始")
					ctx :=context.Background()
					params := svr.params
					params["request_id"] = id
					resp, err := svr.handler.Handle(ctx,params)
					if err != nil{
						it.Logger.Info(err.Error())
					}else{
						it.Logger.Info(resp)
					}
					log.Info(name + ":任务结束")
				}

			}
		}(name,svr,ticker,errChans[name])
	}
	for _,errChan := range errChans{
		e := <- errChan
		if e != nil{
			return e
		}
	}
	return nil
}