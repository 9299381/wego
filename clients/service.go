package clients

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

// 为统一php模式而封装
// micro -> service,  service ->route
func Service(service string) *microService {
	return &microService{
		service: service,
		params:  make(map[string]interface{}),
	}
}

type microService struct {
	service string
	api     string
	params  map[string]interface{}
}

func (s *microService) Api(api string) *microService {
	s.api = api
	return s
}

func (s *microService) Params(params map[string]interface{}) *microService {
	s.params = params
	return s
}

func (s *microService) Run() (resp contracts.Response) {
	entity, err := GetConsulService(s.service)
	if err != nil {
		resp = contracts.ResponseFailed(err)
		return
	}
	tag := entity.Service.Tags[0]
	host := fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
	if tag == "http" {
		return NewHttpPostCall(host, s.api, s.params)
	} else if tag == "grpc" {
		resp = NewGrpcCall(host, s.api, s.params)
	}
	return
}
