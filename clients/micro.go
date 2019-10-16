package clients

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

// 为统一php模式而封装
// micro -> service,  service ->route
func Micro(micro string) *microService {
	return &microService{
		micro:  micro,
		params: make(map[string]interface{}),
	}
}

type microService struct {
	micro   string
	service string
	params  map[string]interface{}
}

func (s *microService) Service(service string) *microService {
	s.service = service
	return s
}

func (s *microService) Params(params map[string]interface{}) *microService {
	s.params = params
	return s
}

func (s *microService) Run() (resp contracts.Response) {
	entity, err := GetConsulService(s.micro)
	if err != nil {
		resp = contracts.ResponseFailed(err)
		return
	}
	tag := entity.Service.Tags[0]
	host := fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
	if tag == "http" {
		return NewHttpPostCall(host, s.service, s.params)
	} else if tag == "grpc" {
		resp = NewGrpcCall(host, s.service, s.params)
	}
	return
}
