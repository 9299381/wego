package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
)

type SqlController struct {
}

func (s *SqlController) Handle(ctx contracts.Context) (interface{}, error) {
	_ = services.Pipe().Middle(&service.SqlService{}).Line(ctx)
	ret := ctx.Get("user")
	return ret, nil

}
