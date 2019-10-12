package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/dto"
	"github.com/9299381/wego/tools/convert"
	"github.com/9299381/wego/validations"
)

type ValidController struct {
}

func (it *ValidController) Handle(ctx contracts.Context) (interface{}, error) {

	return nil, nil
}

func (it *ValidController) Valid(ctx contracts.Context) error {
	req := ctx.Request()
	st := &dto.TestDto{}
	err := convert.Map2Struct(req, st)
	if err != nil {
		return err
	}
	err = validations.Valid(st)
	if err != nil {
		return err
	}
	return nil
}
