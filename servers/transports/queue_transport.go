package transports

import (
	"github.com/9299381/wego/servers/codecs"
	"github.com/9299381/wego/servers/commons"
	"github.com/go-kit/kit/endpoint"
)

func NewQueue(endpoint endpoint.Endpoint) *commons.Server {
	return commons.NewServer(
		endpoint,
		codecs.QueueServerDecodeRequest,
		codecs.QueueServerEncodeResponse,
	)
}
