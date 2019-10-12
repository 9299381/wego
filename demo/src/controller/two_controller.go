package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
	"github.com/9299381/wego/tools/idwork"
)

type TwoController struct {
}

func (it *TwoController) Handle(ctx contracts.Context) (interface{}, error) {

	chain := services.Chain(
		&service.TwoService{},
	)
	_ = chain.Handle(ctx)

	ret := &TwoResponse{
		Id:       idwork.ID(),
		UserName: ctx.GetValue("one").(string),
	}
	return ret, nil
}

func (it *TwoController) Valid(ctx contracts.Context) error {
	return nil
}

// swagger:parameters FirstController
type TwoRequest struct {
	Param1 string `json:"param_1"`
	Param2 int    `json:"param_2"`
}

// swagger:response FirstController
type TwoResponse struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}
