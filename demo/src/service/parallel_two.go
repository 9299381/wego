package service

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type ParallelTwo struct {
}

func (it *ParallelTwo) Handle(ctx contracts.Context) error {
	ctx.Log.Info(ctx.GetValue("controller"))
	fmt.Println("two~~~~~~~~~~~~")
	ctx.SetValue("aaa", "ccc")
	ctx.SetValue("two", "two")

	return nil
}
