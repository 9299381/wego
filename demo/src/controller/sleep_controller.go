package controller

import (
	"github.com/9299381/wego/contracts"
	"time"
)

//用于测试并行,串行
type SleepController struct {
}

func (it *SleepController) Handle(ctx contracts.Context) (interface{}, error) {

	time.Sleep(10 * time.Second)
	ctx.Log.Info("sleep ......")

	return nil, nil
}
func (it *SleepController) Valid(ctx contracts.Context) error {
	return nil
}
