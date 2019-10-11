package mqtts

import (
	"context"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/commons"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Server struct {
	topics   map[string]*commons.CommHandler
	Logger   contracts.ILogger
	Parallel bool //并行处理
}

func NewServer() *Server {
	config := (&configs.MqttConfig{}).Load()
	ss := &Server{
		topics:   make(map[string]*commons.CommHandler),
		Parallel: config.Parallel,
	}

	return ss
}
func (it *Server) Register(name string, handler *commons.CommHandler) {
	it.topics[name] = handler

}

func (it *Server) Serve() error {
	errChans := make(map[string]chan error)
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
		go it.worker(topic, handler, errChans[topic])
	}

}
func (it *Server) worker(t string, h *commons.CommHandler, e chan error) {
	token := mc.Subscribe(t, 0, func(
		client mqtt.Client, message mqtt.Message) {
		if it.Parallel {
			go it.process(h, message)
		} else {
			it.process(h, message)
		}
	})
	if token.Wait() && token.Error() != nil {
		e <- token.Error()
	}
}
func (it *Server) process(h *commons.CommHandler, Message mqtt.Message) {
	it.Logger.Info("subscribe topic:", Message.Topic())
	resp, err := h.Handle(context.Background(), Message.Payload())
	if err != nil {
		it.Logger.Error(err)

	} else {
		it.Logger.Info(resp)
	}
}
