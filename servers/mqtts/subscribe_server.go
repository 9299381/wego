package mqtts

import (
	"context"
	"errors"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/constants"
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
	if GetIns() != nil {
		errChans := make(map[string]chan error)
		it.work(errChans)
		for _, errChan := range errChans {
			if errChan != nil {
				it.Logger.Info(<-errChan)
			}
		}
	} else {
		it.Logger.Info(errors.New(constants.ErrMQTTConnect))
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
	it.Logger.Infof("Subscribe topic:%s", t)
	token := GetIns().Subscribe(t, 0, func(
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

func (it *Server) Close() {
	if GetIns() != nil {
		for topic := range it.topics {
			GetIns().Unsubscribe(topic)
			it.Logger.Infof("Unsubscribe topic:%s", topic)
		}
		GetIns().Disconnect(250)
	}
}
