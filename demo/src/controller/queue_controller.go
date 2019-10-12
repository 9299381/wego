package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/queues"
)

type QueueController struct {
}

func (it *QueueController) Handle(ctx contracts.Context) (interface{}, error) {

	msg := make(map[string]interface{})
	msg["aaa"] = "bbb"

	err := queues.Fire("demo1", "queue_test", msg)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (it *QueueController) Valid(ctx contracts.Context) error {
	return nil
}
