package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/dto"
)

type ValidController struct {
	*contracts.Controller
}

func (s *ValidController) Handle(ctx contracts.Context) (interface{}, error) {

	return nil, nil
}

func (s *ValidController) GetRules() interface{} {
	return &dto.TestDto{}
}
