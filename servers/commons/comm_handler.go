package commons

import (
	"context"
)

type CommHandler struct {
	Handler Handler
}

func (it *CommHandler) Handle(ctx context.Context, req interface{}) (interface{}, error) {
	rsp, err := it.Handler.ServeHandle(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp, err
}

//该接口的实现是为了 cronjob
func (it *CommHandler) Run() {
	ctx := context.Background()
	req := make(map[string]interface{})
	_, _ = it.Handler.ServeHandle(ctx, req)
}
