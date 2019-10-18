package service

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type ParallelOne struct {
}

func (s *ParallelOne) Handle(ctx contracts.Context) error {
	ctx.Log.Info(ctx.Get("controller"))
	fmt.Println("one~~~~~~~~~~~~")
	ctx.Set("aaa", "bbb")
	ctx.Set("one", "one")
	//return errors.New("error")
	return nil
}
