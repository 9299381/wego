package mqtts

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/commons"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Server struct {
	topics map[string]*commons.CommHandler
	Logger contracts.ILogger
}

func NewServer() *Server {
	ss := &Server{
		topics: make(map[string]*commons.CommHandler),
	}

	return ss
}
func (it *Server) Register(name string, handler *commons.CommHandler) {
	it.topics[name] = handler

}

func (it *Server) Serve() error {
	errChans := make(map[string]chan error)
	//errChans["connectError"] = make(chan error)
	//if token := mc.Connect(); token.Wait() && token.Error() != nil {
	//	errChans["connectError"] <- token.Error()
	//}
	it.work(errChans)
	for _, errChan := range errChans {
		if errChan != nil {
			it.Logger.Info(<-errChan)
		}
	}
	return nil
}
func (it *Server) work(errChans map[string]chan error) {
	it.Logger.Info("MQTT Subscribe Server Start")
	for topic, handler := range it.topics {
		errChans[topic] = make(chan error)
		go func(t string, h *commons.CommHandler, e chan error) {
			token := mc.Subscribe(t, 0, func(
				client mqtt.Client, message mqtt.Message) {
				it.Logger.Info("subscribe topic:", message.Topic())
				resp, err := h.Handle(context.Background(), message.Payload())
				if err != nil {
					it.Logger.Error(err)

				} else {
					it.Logger.Info(resp)
				}
			})
			if token.Wait() && token.Error() != nil {
				e <- token.Error()
			}

		}(topic, handler, errChans[topic])
	}

}
