package service

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/queues"
)

type TestQueue struct {
	repo string
	next contracts.IService
}

func (it *TestQueue) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *TestQueue) Handle(ctx contracts.Context) error {

	msg := make(map[string]interface{})
	msg["aaa"] = "bbb"

	err := queues.Fire("demo1", "queue_test", msg)
	if err != nil {
		return err
	}

	return it.next.Handle(ctx)
}
