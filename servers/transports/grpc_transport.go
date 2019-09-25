package transports

import (
	"github.com/9299381/wego/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
	GrpcTransport "github.com/go-kit/kit/transport/grpc"
)

func NewGRPC(endpoint endpoint.Endpoint) *GrpcTransport.Server {
	return GrpcTransport.NewServer(
		endpoint,
		codecs.GprcDecodeRequest,
		codecs.GprcEncodeResponse,
	)
}
