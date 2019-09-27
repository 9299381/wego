package clients

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type microService struct {
	micro   string
	service string
	params  map[string]interface{}
}

func (it *microService) Service(service string) *microService {
	it.service = service
	return it
}

func (it *microService) Params(params map[string]interface{}) *microService {
	it.params = params
	return it
}

func (it *microService) Run() (resp contracts.Response) {
	entity, err := GetConsulService(it.micro)
	if err != nil {
		resp = contracts.ResponseFailed(err)
	}
	tag := entity.Service.Tags[0]
	host := fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
	if tag == "http" {
		return NewHttpPostCall(host, it.service, it.params)
	} else if tag == "grpc" {
		resp = NewGrpcCall(host, it.service, it.params)
	}
	return
}
