package service

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type ParallelOne struct {
}

func (it *ParallelOne) Handle(ctx contracts.Context) error {
	ctx.Log.Info(ctx.GetValue("controller"))
	fmt.Println("one~~~~~~~~~~~~")
	ctx.SetValue("aaa", "bbb")
	ctx.SetValue("one", "one")
	//return errors.New("error")
	return nil
}
