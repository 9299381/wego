package events

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/idwork"
	"time"
)

type Server struct {
	Ticker *time.Ticker
	Logger contracts.ILogger
}

func NewServer() *Server {
	ss := &Server{}
	return ss
}

func (it *Server) Serve() error {
	errChan := make(chan error)
	events := it.getEvent()
	it.runEvent(events, errChan)
	err := <-errChan
	if err != nil {
		it.Logger.Info(err)
	}
	return nil
}

// 每3秒检查一次任务
func (it *Server) getEvent() <-chan *contracts.Payload {
	events := make(chan *contracts.Payload)
	go func() {
		for {
			select {
			case <-it.Ticker.C:
				event := Event.Pop()
				if event != nil {
					select {
					case events <- event:
					}
				}
			}
		}
	}()
	return events
}
func (it *Server) runEvent(events <-chan *contracts.Payload, errChan chan error) {
	go func(errChan chan error) {
		for event := range events {
			filter, ok := Handlers[event.Route]
			if ok {
				ctx := context.Background()
				id := idwork.ID()
				request := contracts.Request{
					Id:   id,
					Data: event.Params,
				}
				resp, err := filter(ctx, request)
				if err != nil {
					errChan <- err
				} else {
					it.Logger.Info("事件结果:", resp)
				}
			}
		}
	}(errChan)
}
