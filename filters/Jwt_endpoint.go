package filters

import (
	"context"
	"errors"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/convert"
	"github.com/9299381/wego/tools/jwt"
	"github.com/go-kit/kit/endpoint"
)

type JwtEndpoint struct {
	next endpoint.Endpoint
}

func (s *JwtEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	s.next = next
	return s
}

func (s *JwtEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(contracts.Request)
		token := req.Data["authToken"]
		if token == nil || token == "" {
			return nil, errors.New(constants.ErrNoToken)
		}
		claim, err := jwt.New().VerifyToken(token.(string))
		if err != nil {
			return nil, err
		}
		req.Data["claim"] = convert.Struct2Map(claim)
		//这里进行token的jwt认证
		return s.next(ctx, req)
	}
}
