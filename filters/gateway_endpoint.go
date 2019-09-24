package filters

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
)

type GateWayEndpoint struct {
}

func (it *GateWayEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	return it
}

func (it *GateWayEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return contracts.ResponseSucess("GATEWAY"), nil
	}
}
