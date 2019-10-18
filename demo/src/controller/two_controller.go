package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
	"github.com/9299381/wego/tools/idwork"
)

type TwoController struct {
	*contracts.Controller
}

// swagger:route Get /demo/two 分组1 twoController
// Test swagger
// This will .......
//     Responses:
//       200: twoResponse

func (s *TwoController) Handle(ctx contracts.Context) (interface{}, error) {
	_ = services.Pipe().
		Middle(&service.TwoService{}).
		Line(ctx)
	ret := &TwoResponse{
		Id:       idwork.ID(),
		UserName: ctx.Get("one").(string),
	}
	return ret, nil
}

// swagger:parameters twoController
type TwoRequest struct {
}

// swagger:response twoResponse
type TwoResponse struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}
