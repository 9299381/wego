package transports

import (
	"github.com/9299381/wego/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
	HttpTransport "github.com/go-kit/kit/transport/http"
)

func NewWeb(endpoint endpoint.Endpoint) *HttpTransport.Server {
	return HttpTransport.NewServer(
		endpoint,
		codecs.WebDecodeRequest,
		codecs.WebEncodeResponse,
	)
}
