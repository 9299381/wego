package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
)

type ParallelController struct {
}

func (it *ParallelController) Handle(ctx contracts.Context) (interface{}, error) {
	ctx.SetValue("controller", "parallel")

	//并行service中间件
	err := services.New().
		Middle(&service.ParallelOne{}).
		Middle(&service.ParallelTwo{}).
		Parallel(ctx)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m["aaa"] = ctx.GetValue("aaa")               //不确定值,
	m["one"] = ctx.GetValue("one")               //one中设置
	m["two"] = ctx.GetValue("two")               //two中设置
	m["controller"] = ctx.GetValue("controller") //controller中设置

	return m, nil
}
func (it *ParallelController) Valid(ctx contracts.Context) error {
	return nil
}
