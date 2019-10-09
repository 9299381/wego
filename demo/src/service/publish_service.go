package service

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/mqtts"
)

type Publish struct {
	next contracts.IService
}

func (it *Publish) Next(srv contracts.IService) contracts.IService {
	//这里可以做些事情
	it.next = srv
	return it
}

func (it *Publish) Handle(ctx contracts.Context) error {

	m := make(map[string]interface{})
	m["pub"] = "pub"
	m["sub"] = "sub"
	err := mqtts.Publish("sub_test", m)
	if err != nil {
		return err
	}
	ctx.Response("aa", "bb")
	return it.next.Handle(ctx)
}
