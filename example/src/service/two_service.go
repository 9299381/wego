package service

import (
	"github.com/9299381/wego/contracts"
)

type TwoService struct {
	repo string
	next contracts.IService
}

func (it *TwoService)Next(srv contracts.IService) contracts.IService  {
	//这里可以做些事情
	it.repo = "aaa-bbb"
	it.next = srv
	return it
}

func (it *TwoService)Handle(ctx contracts.Context) error {

	ctx.Response("one","tow")
	//return errors.New("8888","系统错误")

	ctx.Log.Info("two......")
	return it.next.Handle(ctx)
}