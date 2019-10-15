package service

import (
	"github.com/9299381/wego/contracts"
)

type OneService struct {
}

func (it *OneService) Handle(ctx contracts.Context) error {
	ctx.Log.Info("one....")
	ctx.SetValue("k.a", "a")
	ctx.SetValue("k.b", "b")
	ctx.SetValue("k.a", "b")

	ctx.Log.Info(ctx.GetValue("k"))
	return nil
}
