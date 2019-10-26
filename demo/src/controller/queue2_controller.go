package controller

import (
	"fmt"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/dto"
)

type Queue2Controller struct {
}

func (s *Queue2Controller) Handle(ctx contracts.Context) (interface{}, error) {
	req := ctx.Request().(*dto.TaskRequest)
	for _, v := range req.OrderList {
		fmt.Print(v)
	}
	return nil, nil
}
func (s *Queue2Controller) GetRules() interface{} {
	return &dto.TaskRequest{}
}
