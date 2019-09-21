package service

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/dto"
	"github.com/9299381/wego/tools/convert"
	"github.com/9299381/wego/validations"
)

type ValidService struct {
	repo string
	next contracts.IService
}

func (it *ValidService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}
func (it *ValidService) Handle(ctx contracts.Context) error {

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
	return it.next.Handle(ctx)
}
