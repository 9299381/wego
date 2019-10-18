package service

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type ParallelTwo struct {
}

func (s *ParallelTwo) Handle(ctx contracts.Context) error {
	ctx.Log.Info(ctx.Get("controller"))
	fmt.Println("two~~~~~~~~~~~~")
	ctx.Set("aaa", "ccc")
	ctx.Set("two", "two")

	return nil
}
