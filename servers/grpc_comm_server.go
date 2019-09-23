package servers

import (
	"context"
	"github.com/9299381/wego"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/servers/transports"
	"github.com/9299381/wego/servers/transports/protobuf"
	"github.com/go-kit/kit/endpoint"
	GrpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"net"
)

type GrpcCommServer struct {
	*grpc.Server
}

func NewGrpcCommServer() *GrpcCommServer {
	return &GrpcCommServer{
		Server: grpc.NewServer(),
	}
}
func (it *GrpcCommServer) Route(name string, endpoint endpoint.Endpoint) {

	sd := it.getServiceDesc(name)
	service := &grpcService{
		handler: transports.NewGRPC(endpoint),
	}
	it.RegisterService(&sd, service)
}

func (it *GrpcCommServer) Load() {

	//注册通用路由
}

func (it *GrpcCommServer) Start() error {
	config := (&configs.GrpcConfig{}).Load().(*configs.GrpcConfig)
	wego.App.Logger.Info("Grpc Server Start ", config.GrpcPort)
	lis, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		return err
	}
	return it.Serve(lis)
}

//--------------

func (it *GrpcCommServer) getServiceDesc(name string) grpc.ServiceDesc {
	var serviceDesc = grpc.ServiceDesc{
		ServiceName: "protobuf." + name,
		HandlerType: (*protobuf.ServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Handle",
				Handler:    it.serviceHandleHandler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "message.proto",
	}
	return serviceDesc
}

func (it *GrpcCommServer) serviceHandleHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	//interceptor 拦截器
	if interceptor == nil {
		return srv.(protobuf.ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(protobuf.ServiceServer).Handle(ctx, req.(*protobuf.Request))
	}
	return interceptor(ctx, in, info, handler)
}

type grpcService struct {
	handler GrpcTransport.Handler
}

func (it *grpcService) Handle(ctx context.Context, req *protobuf.Request) (*protobuf.Response, error) {
	_, rsp, err := it.handler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*protobuf.Response), err
}
