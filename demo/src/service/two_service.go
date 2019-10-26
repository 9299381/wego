package service

import (
	"github.com/9299381/wego/contracts"
)

type TwoService struct {
}

func (s *TwoService) Handle(ctx contracts.Context) error {

	ctx.Set("one", "tow")
	ctx.Log.Info("two......")
	return nil
}
