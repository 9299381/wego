package controller

import (
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
)

type ConsulController struct {
	*contracts.Controller
}

func (s *ConsulController) Handle(ctx contracts.Context) (interface{}, error) {

	entity, _ := clients.GetConsulService(args.Name)
	ctx.Log.Info(entity.Service.Service)
	ctx.Log.Info(entity.Service.Address)
	ctx.Log.Info(entity.Service.Port)
	tag := entity.Service.Tags[0]
	ctx.Log.Info(tag)

	return nil, nil
}
