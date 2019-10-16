package filters

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
)

type GateWayEndpoint struct {
	next endpoint.Endpoint
}

func (s *GateWayEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	s.next = next
	return s
}

func (s *GateWayEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(contracts.Request)
		req.Data["GATEWAY"] = "GATEWAY"
		if s.next == nil {
			response = contracts.ResponseSucess(req.Data)
		} else {
			response, err = s.next(ctx, req)
			res := response.(contracts.Response)
			if res.Code == "0000" {
				m, b := res.Data.(map[string]interface{})
				if b && m != nil {
					for k, v := range m {
						req.Data[k] = v
					}
					response = contracts.ResponseSucess(req.Data)
				}
			}
		}
		return
	}
}
