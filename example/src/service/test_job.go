package service

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type TestJob struct {
	repo string
	next contracts.IService
}

func (it *TestJob)Next(srv contracts.IService) contracts.IService  {
	it.next = srv
	return it
}

func (it *TestJob)Handle(ctx contracts.Context) error {

	msg := make(map[string]interface{})
	msg["aaa"] = "bbb"

	err := wego.Queue("demo1","queue_test",msg)
	if err !=nil{
		return err
	}

	return it.next.Handle(ctx)
}