package servers

import (
	"context"
	"github.com/9299381/wego"
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
	ss.Logger = loggers.GetLog()
	return ss
}
func (s *GrpcCommServer) Route(name string, endpoint endpoint.Endpoint) {

	sd := s.getServiceDesc(name)
	service := &grpcService{
		handler: transports.NewGRPC(endpoint),
	}
	s.RegisterService(&sd, service)
}

func (s *GrpcCommServer) Load() {

	//注册通用路由
	grpc_health_v1.RegisterHealthServer(s.Server, &HealthImpl{})
}

func (s *GrpcCommServer) Start() error {
	config := configs.LoadGrpcConfig()
	address := config.GrpcHost + ":" + config.GrpcPort
	s.Logger.Info("Grpc Server Start ", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

//--------------

func (s *GrpcCommServer) getServiceDesc(name string) grpc.ServiceDesc {
	var serviceDesc = grpc.ServiceDesc{
		ServiceName: "protobuf." + name,
		HandlerType: (*protobuf.ServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Handle",
				Handler:    s.serviceHandleHandler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "message.proto",
	}
	return serviceDesc
}

func (s *GrpcCommServer) serviceHandleHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
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

func (s *GrpcCommServer) Close() {
	v, ok := wego.App.Consul["grpc"]
	if ok {
		v.Deregister()
	}
}

type grpcService struct {
	handler GrpcTransport.Handler
}

func (s *grpcService) Handle(ctx context.Context, req *protobuf.Request) (*protobuf.Response, error) {
	_, rsp, err := s.handler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*protobuf.Response), err
}

type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，
// 这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (s *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
func (s *HealthImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {

	return nil
}
