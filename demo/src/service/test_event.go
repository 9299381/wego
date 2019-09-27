package service

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/events"
)

type TestEvent struct {
	repo string
	next contracts.IService
}

func (it *TestEvent) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *TestEvent) Handle(ctx contracts.Context) error {

	params := map[string]interface{}{}
	payload := &contracts.Payload{
		Route:  "two",
		Params: params,
	}
	events.Fire(payload)

	return it.next.Handle(ctx)
}
