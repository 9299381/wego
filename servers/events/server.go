package events

import (
	"context"
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"time"
)

type Server struct {
	Ticker *time.Ticker
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
		wego.App.Logger.Info(err)
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
				event := wego.App.Event.Pop()
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
			endpoint := wego.Handler(event.Route)
			if endpoint != nil {
				ctx := context.Background()
				id := wego.ID()
				request := contracts.Request{
					Id:   id,
					Data: event.Params,
				}
				resp, err := endpoint(ctx, request)
				if err != nil {
					errChan <- err
				} else {
					wego.App.Logger.Info("事件结果:", resp)
				}
			}
		}
	}(errChan)
}
