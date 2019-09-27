package service

import (
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
)

type ConsulService struct {
	next contracts.IService
}

func (it *ConsulService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *ConsulService) Handle(ctx contracts.Context) error {

	entity, _ := clients.GetConsulService(args.Name)
	ctx.Log.Info(entity.Service.Service)
	ctx.Log.Info(entity.Service.Address)
	ctx.Log.Info(entity.Service.Port)
	tag := entity.Service.Tags[0]
	ctx.Log.Info(tag)

	return it.next.Handle(ctx)

}
