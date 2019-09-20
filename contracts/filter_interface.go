package contracts

import "github.com/go-kit/kit/endpoint"

type IFilter interface {
	Next(next endpoint.Endpoint) IFilter
	Make() endpoint.Endpoint
}
