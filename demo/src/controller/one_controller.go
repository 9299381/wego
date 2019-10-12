// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
	"github.com/9299381/wego/tools/idwork"
)

type OneController struct {
}

//swagger:route GET /demo/one  users FirstController
func (it *OneController) Handle(ctx contracts.Context) (interface{}, error) {

	err := services.Chain(
		&service.OneService{},
		&service.TwoService{},
	).Handle(ctx)

	if err != nil {
		return nil, err
	}
	ret := &OneResponse{
		Id:       idwork.ID(),
		UserName: ctx.GetValue("k.a").(string),
	}
	return ret, nil
}

func (it *OneController) Valid(ctx contracts.Context) error {
	return nil
}

// swagger:parameters FirstController
type OneRequest struct {
	Param1 string `json:"param_1"`
	Param2 int    `json:"param_2"`
}

// swagger:response FirstController
type OneResponse struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}
