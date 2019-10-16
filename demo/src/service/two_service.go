package service

import (
	"github.com/9299381/wego/contracts"
)

type TwoService struct {
}

func (s *TwoService) Handle(ctx contracts.Context) error {

	ctx.SetValue("one", "tow")
	//if ctx.Request("a").(string) == "a" {
	//	panic(errors.New("9988::严重错误"))
	//}
	ctx.Log.Info("two......")
	return nil
}
