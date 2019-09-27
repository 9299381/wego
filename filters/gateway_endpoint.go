package filters

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
)

type GateWayEndpoint struct {
	next endpoint.Endpoint
}

func (it *GateWayEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	it.next = next
	return it
}

func (it *GateWayEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req := request.(contracts.Request)
		//req.Data["GATEWAY"] = "GATEWAY"
		//return contracts.ResponseSucess(req.Data), nil
		return contracts.ResponseSucess("GATEWAY"), nil

	}
}
