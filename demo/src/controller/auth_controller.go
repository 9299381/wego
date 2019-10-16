package controller

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type AuthController struct {
	*contracts.Controller
}

func (s *AuthController) Handle(ctx contracts.Context) (interface{}, error) {

	fmt.Println(ctx.GetValue("request.claim.Id"))
	fmt.Println(ctx.GetValue("request.claim.Name"))

	return nil, nil
}
