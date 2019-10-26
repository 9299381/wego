package controller

import (
	"fmt"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
)

type OneController struct {
	*contracts.Controller
}

// ServeAPI serves the API for this record store
func (s *OneController) Handle(ctx contracts.Context) (interface{}, error) {
	// swagger:route Get /demo/one 分组1 oneController
	// Test swagger
	// This will .......
	//     Responses:
	//       200: oneResponse
	req := ctx.Request().(*oneRequest)
	fmt.Println(req)
	fmt.Println(ctx.Get("request"))

	err := services.Pipe().
		Middle(&service.OneService{}).
		Middle(&service.TwoService{}).
		Line(ctx)
	if err != nil {
		return nil, err
	}
	ret := &oneResponse{
		Id:       ctx.Get("request.id").(string),
		UserName: req.Param1,
		Age:      req.Param2,
	}
	return ret, nil
}

func (s *OneController) GetRules() interface{} {
	return &oneRequest{}
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
	// 整形
	Age int `json:"age"`
}
