package filters

import (
	"context"
	"errors"
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
	"github.com/juju/ratelimit"
	"time"
)

type LimitEndpoint struct {
	next endpoint.Endpoint
}

func (s *LimitEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	s.next = next
	return s
}

func (s *LimitEndpoint) Make() endpoint.Endpoint {
	limit := ratelimit.NewBucket(time.Second*1, 3)
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if limit.TakeAvailable(1) == 0 {
			return nil, errors.New("Rate limit exceed!")
		}
		return s.next(ctx, request)
	}
}
