package service

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type OneService struct {
	repo string
	next contracts.IService
}

func (it *OneService)Next(srv contracts.IService) contracts.IService  {
	//这里可以做些事情
	it.repo = "aaa-bbb"
	it.next = srv
	return it
}

func (it *OneService)Handle(ctx contracts.Context) error {
	ctx.Response("one","one")


	ctx.SetValue("k.a","a")
	ctx.SetValue("k.b","b")
	ctx.SetValue("k.a","b")

	fmt.Println(ctx.GetValue("k"))
	ctx.Log.Info("one.....")

	return it.next.Handle(ctx)
}