package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/events"
)

type EventController struct {
}

func (it *EventController) Handle(ctx contracts.Context) (interface{}, error) {
	params := make(map[string]interface{})
	payload := &contracts.Payload{
		Route:  "two",
		Params: params,
	}
	events.Fire(payload)
	return nil, nil
}

func (it *EventController) Valid(ctx contracts.Context) error {
	return nil
}
