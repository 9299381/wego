package filters

import (
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
)

func Chain(endpoints ...contracts.IFilter) endpoint.Endpoint {
	len := len(endpoints) - 1
	for i := 0; i < len; i++ {
		endpoints[i].Next(endpoints[i+1].Make())
	}
	return endpoints[0].Make()
}
