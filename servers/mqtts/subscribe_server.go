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
	topics       map[string]*commons.CommHandler
	Logger       contracts.ILogger
	Parallel     bool //并行处理
	SubscribeQos byte
}

func NewServer() *Server {
	config := configs.LoadMqttConfig()
	ss := &Server{
		topics:       make(map[string]*commons.CommHandler),
		Parallel:     config.Parallel,
		SubscribeQos: config.SubscribeQos,
	}
	return ss
}
func (s *Server) Register(name string, handler *commons.CommHandler) {
	s.topics[name] = handler

}
func (s *Server) Serve() error {
	if GetIns() != nil {
		errChans := make(map[string]chan error)
		s.work(errChans)
		for _, errChan := range errChans {
			if errChan != nil {
				s.Logger.Info(<-errChan)
			}
		}
	} else {
		s.Logger.Info(errors.New(constants.ErrMQTTConnect))
	}
	return nil
}

func (s *Server) work(errChans map[string]chan error) {
	s.Logger.Info("MQTT Subscribe Server Start")
	for topic, handler := range s.topics {
		errChans[topic] = make(chan error)
		go s.worker(topic, handler, errChans[topic])
	}

}
func (s *Server) worker(t string, h *commons.CommHandler, e chan error) {
	s.Logger.Infof("Subscribe topic:%s", t)
	token := GetIns().Subscribe(t, s.SubscribeQos, func(
		client mqtt.Client, message mqtt.Message) {
		if s.Parallel {
			go s.process(h, message)
		} else {
			s.process(h, message)
		}
	})
	if token.Wait() && token.Error() != nil {
		e <- token.Error()
	}
}
func (s *Server) process(h *commons.CommHandler, Message mqtt.Message) {
	s.Logger.Info("subscribe topic:", Message.Topic())
	resp, err := h.Handle(context.Background(), Message.Payload())
	if err != nil {
		s.Logger.Error(err)

	} else {
		s.Logger.Info(resp)
	}
}

func (s *Server) Close() {
	if GetIns() != nil {
		for topic := range s.topics {
			GetIns().Unsubscribe(topic)
			s.Logger.Infof("Unsubscribe topic:%s", topic)
		}
		GetIns().Disconnect(250)
	}
}
