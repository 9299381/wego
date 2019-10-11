package service

import (
	"github.com/9299381/wego/contracts"
	"time"
)

type SleepService struct {
	next contracts.IService
}

func (it *SleepService) Next(srv contracts.IService) contracts.IService {
	//这里可以做些事情
	it.next = srv
	return it
}

func (it *SleepService) Handle(ctx contracts.Context) error {

	time.Sleep(10 * time.Second)
	ctx.Log.Info("sleep ......")
	return it.next.Handle(ctx)
}
