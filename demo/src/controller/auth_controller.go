package controller

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type AuthController struct {
}

func (s *AuthController) Handle(ctx contracts.Context) (interface{}, error) {

	fmt.Println(ctx.Get("request.claim.Id"))
	fmt.Println(ctx.Get("request.claim.Name"))

	return nil, nil
}
