package transports

import (
	"github.com/9299381/wego/servers/codecs"
	"github.com/go-kit/kit/endpoint"
	HttpTransport "github.com/go-kit/kit/transport/http"
)

func NewHTTP(endpoint endpoint.Endpoint) *HttpTransport.Server {
	return HttpTransport.NewServer(
		endpoint,
		codecs.HttpFormDecodeRequest,
		codecs.HttpEncodeResponse,
	)
}
