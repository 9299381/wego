package controller

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type AuthController struct {
}

func (it *AuthController) Handle(ctx contracts.Context) (interface{}, error) {

	fmt.Println(ctx.GetValue("request.claim.Id"))
	fmt.Println(ctx.GetValue("request.claim.Name"))

	return nil, nil
}

func (it *AuthController) Valid(ctx contracts.Context) error {
	return nil
}
