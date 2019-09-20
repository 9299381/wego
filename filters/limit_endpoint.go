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

func (it *LimitEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	it.next = next
	return it
}

func (it *LimitEndpoint) Make() endpoint.Endpoint {
	limit := ratelimit.NewBucket(time.Second*1, 3)
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if limit.TakeAvailable(1) == 0 {
			return nil, errors.New("Rate limit exceed!")
		}
		return it.next(ctx, request)
	}
}
