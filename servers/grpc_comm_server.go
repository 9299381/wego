package servers

import (
	"context"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/transports"
	"github.com/9299381/wego/servers/transports/protobuf"
	"github.com/go-kit/kit/endpoint"
	GrpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

type GrpcCommServer struct {
	*grpc.Server
	Logger contracts.ILogger
}

func NewGrpcCommServer() *GrpcCommServer {
	ss := &GrpcCommServer{
		Server: grpc.NewServer(),
	}
	ss.Logger = loggers.Log
	return ss
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
	grpc_health_v1.RegisterHealthServer(it.Server, &HealthImpl{})
}

func (it *GrpcCommServer) Start() error {
	config := (&configs.GrpcConfig{}).Load()
	address := config.GrpcHost + ":" + config.GrpcPort
	it.Logger.Info("Grpc Server Start ", address)
	lis, err := net.Listen("tcp", address)
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

type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，
// 这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (it *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
func (it *HealthImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {

	return nil
}
