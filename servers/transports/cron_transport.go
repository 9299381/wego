package transports

import (
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
)

func NewCronJob(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.CronDecodeRequest,
		codecs.CronEncodeResponse,
	)
}
