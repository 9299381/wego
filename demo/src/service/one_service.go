package service

import (
	"github.com/9299381/wego/contracts"
)

type OneService struct {
}

func (s *OneService) Handle(ctx contracts.Context) error {
	ctx.Log.Info("one....")
	ctx.Set("k.a", "a")
	ctx.Set("k.b", "b")
	ctx.Set("k.a", "b")

	ctx.Log.Info(ctx.Get("k"))
	return nil
}
