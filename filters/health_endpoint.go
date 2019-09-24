package filters

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
)

type HealthEndpoint struct {
}

func (it *HealthEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	return it
}

func (it *HealthEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return contracts.ResponseSucess("SERVING"), nil
	}
}
