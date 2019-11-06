package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
)

type ParallelController struct {
}

func (s *ParallelController) Handle(ctx contracts.Context) (interface{}, error) {
	ctx.Set("controller", "parallel")

	//并行service中间件
	err := services.Pipe().
		Middle(&service.ParallelOne{}).
		Middle(&service.ParallelTwo{}).
		Parallel(ctx)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m["aaa"] = ctx.Get("aaa")               //不确定值,
	m["one"] = ctx.Get("one")               //one中设置
	m["two"] = ctx.Get("two")               //two中设置
	m["controller"] = ctx.Get("controller") //controller中设置

	return m, nil
}
