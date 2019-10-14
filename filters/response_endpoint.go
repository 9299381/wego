package filters

import (
	"context"
	"errors"
	"github.com/9299381/wego"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/loggers"
	"github.com/go-kit/kit/endpoint"
)

type ResponseEndpoint struct {
	next endpoint.Endpoint
}

func (it *ResponseEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	it.next = next
	return it
}

func (it *ResponseEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//全局扑捉错误
		defer func() {
			if err := recover(); err != nil {
				loggers.GetLog().Info(err)
				response = contracts.MakeResponse(nil, err.(error))
			}
		}()
		if wego.App.Status == false {
			err := errors.New(constants.ErrStop)
			return contracts.MakeResponse(nil, err), nil
		}
		response, err = it.next(ctx, request)
		return contracts.MakeResponse(response, err), nil
	}
}
