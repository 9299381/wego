package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/convert"
	"github.com/9299381/wego/tools/idwork"
)

type PostController struct {
}

// swagger:route Post /demo/post 分组2 postController
// Test swagger
// This will .......
//     Responses:
//       200: postResponse

func (it *PostController) Handle(ctx contracts.Context) (interface{}, error) {

	request := &postRequest{}
	err := convert.Map2Struct(ctx.Request(), request)

	if err != nil {
		return nil, err
	}
	ctx.Log.Info(request)

	ret := &postResponse{
		Id:   idwork.ID(),
		Resp: *request,
	}
	return ret, nil
}

func (it *PostController) Valid(ctx contracts.Context) error {
	return nil
}

// swagger:parameters postController
type postRequest struct {
	Id     string `json:"id"`
	Value  string `json:"value"`
	Number int    `json:"number"`
}

// swagger:response postResponse
type postResponse struct {
	Id   string      `json:"id"`
	Resp postRequest `json:"resp"`
}
