package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/dto"
)

type ValidController struct {
}

func (it *ValidController) Handle(ctx contracts.Context) (interface{}, error) {

	return nil, nil
}

func (it *ValidController) GetRules() interface{} {
	return &dto.TestDto{}
}
