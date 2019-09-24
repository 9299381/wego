package service

import (
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
	stdconsul "github.com/hashicorp/consul/api"
	"math/rand"
)

type ConsulService struct {
	next contracts.IService
}

func (it *ConsulService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *ConsulService) Handle(ctx contracts.Context) error {

	client := clients.GetConsullClient()
	entitys, meta, err := client.Service(args.Name, "", false, &stdconsul.QueryOptions{})
	if err != nil {
		panic(err)
	}
	entity := entitys[rand.Int()%len(entitys)]

	ctx.Log.Info(entity.Service.Service)
	ctx.Log.Info(entity.Service.Address)
	ctx.Log.Info(entity.Service.Port)
	tag := entity.Service.Tags[rand.Int()%len(entity.Service.Tags)]
	ctx.Log.Info(tag)
	ctx.Log.Info(meta)

	return it.next.Handle(ctx)

}
