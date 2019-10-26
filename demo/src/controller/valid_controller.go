package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/dto"
)

type ValidController struct {
	*contracts.Controller
}

// swagger:route Post /demo/valid 分组2 validController
// Test swagger
// This will .......
//     Responses:
//       200: postResponse
func (s *ValidController) Handle(ctx contracts.Context) (interface{}, error) {
	st := ctx.Request().(*dto.TestDto)
	resp := &dto.ValidResponse{
		Age:  st.Age,
		List: st.DemoList,
	}
	return resp, nil
}

func (s *ValidController) GetRules() interface{} {
	return &dto.TestDto{}
}
