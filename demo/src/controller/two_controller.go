package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
	"github.com/9299381/wego/tools/idwork"
)

type TwoController struct {
}

// swagger:route Get /demo/two 分组1 twoController
// Test swagger
// This will .......
//     Responses:
//       200: twoResponse

func (it *TwoController) Handle(ctx contracts.Context) (interface{}, error) {
	_ = services.New().
		Middle(&service.TwoService{}).
		Line(ctx)
	ret := &TwoResponse{
		Id:       idwork.ID(),
		UserName: ctx.GetValue("one").(string),
	}
	return ret, nil
}

func (it *TwoController) Valid(ctx contracts.Context) error {
	return nil
}

// swagger:parameters twoController
type TwoRequest struct {
}

// swagger:response twoResponse
type TwoResponse struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}
