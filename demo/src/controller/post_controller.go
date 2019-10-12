package controller

import (
	"github.com/9299381/wego/contracts"
)

type PostController struct {
}

func (it *PostController) Handle(ctx contracts.Context) (interface{}, error) {
	ret := ctx.Request()
	return ret, nil
}

func (it *PostController) Valid(ctx contracts.Context) error {
	return nil
}
