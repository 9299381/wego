package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/mqtts"
)

type PublishController struct {
}

func (s *PublishController) Handle(ctx contracts.Context) (interface{}, error) {

	m := make(map[string]interface{})
	m["pub"] = "pub"
	m["sub"] = "sub"
	err := mqtts.Publish("sub_test", m)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
