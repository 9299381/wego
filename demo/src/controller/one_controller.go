package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
	"github.com/9299381/wego/tools/idwork"
)

type OneController struct {
}

// ServeAPI serves the API for this record store
func (it *OneController) Handle(ctx contracts.Context) (interface{}, error) {
	// swagger:route Get /demo/one 分组1 oneController
	// Test swagger
	// This will .......
	//     Responses:
	//       200: oneResponse
	err := services.Chain(
		&service.OneService{},
		&service.TwoService{},
	).Handle(ctx)
	if err != nil {
		return nil, err
	}
	ret := &oneResponse{
		Id:       idwork.ID(),
		UserName: ctx.GetValue("k.a").(string),
	}
	return ret, nil
}

func (it *OneController) Valid(ctx contracts.Context) error {
	return nil
}

// swagger:parameters oneController
type oneRequest struct {
	// 参数param1的说明
	// 最小1
	// Minimum:1
	// Required:true
	Param1 string `json:"param_1"`
	// 最短4
	// Required:true
	// MinLength:4
	Param2 int `json:"param_2"`
	Param4 int `json:"param_4"`
}

// swagger:response oneResponse
type oneResponse struct {
	Id string `json:"id"`
	// 响应UserName的描述
	UserName string `json:"user_name"`
}
