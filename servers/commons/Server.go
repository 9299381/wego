package commons

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
)

//通用接口
type Handler interface {
	ServeHandle(ctx context.Context, request interface{}) (interface{}, error)
}

//请求
type DecodeRequestFunc func(context.Context, interface{}) (request interface{}, err error)
type EncodeRequestFunc func(context.Context, interface{}) (request interface{}, err error)

//响应
type EncodeResponseFunc func(context.Context, interface{}) (response interface{}, err error)
type DecodeResponseFunc func(context.Context, interface{}) (response interface{}, err error)

type BeforeFunc func(context.Context, interface{}) context.Context
type AfterFunc func(context.Context, interface{}) context.Context

type FinalizerFunc func(ctx context.Context, err error)

type Server struct {
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []BeforeFunc
	after        []AfterFunc
	finalizer    []FinalizerFunc
	errorHandler transport.ErrorHandler
}

type ServerOption func(*Server)

// ServerBefore functions are executed on the gRPC request object before the
// request is decoded.
func ServerBefore(before ...BeforeFunc) ServerOption {
	return func(s *Server) { s.before = append(s.before, before...) }
}

// ServerAfter functions are executed on the gRPC response writer after the
// endpoint is invoked, but before anything is written to the client.
func ServerAfter(after ...AfterFunc) ServerOption {
	return func(s *Server) { s.after = append(s.after, after...) }
}

// ServerErrorLogger is used to log non-terminal errors. By default, no errors
// are logged.
// Deprecated: Use ServerErrorHandler instead.
func ServerErrorLogger(logger log.Logger) ServerOption {
	return func(s *Server) { s.errorHandler = transport.NewLogErrorHandler(logger) }
}

// ServerErrorHandler is used to handle non-terminal errors. By default, non-terminal errors
// are ignored.
func ServerErrorHandler(errorHandler transport.ErrorHandler) ServerOption {
	return func(s *Server) { s.errorHandler = errorHandler }
}

// ServerFinalizer is executed at the end of every gRPC request.
// By default, no finalizer is registered.
func ServerFinalizer(f ...FinalizerFunc) ServerOption {
	return func(s *Server) { s.finalizer = append(s.finalizer, f...) }
}

func NewServer(
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...ServerOption,
) *Server {
	s := &Server{
		e:            e,
		dec:          dec,
		enc:          enc,
		errorHandler: transport.NewLogErrorHandler(log.NewNopLogger()),
	}
	for _, option := range options {
		option(s)
	}
	return s
}

func (s *Server) ServeHandle(ctx context.Context, req interface{}) (resp interface{}, err error) {

	if len(s.finalizer) > 0 {
		defer func() {
			for _, f := range s.finalizer {
				f(ctx, err)
			}
		}()
	}
	for _, f := range s.before {
		ctx = f(ctx, req)
	}

	resp, err = s.dec(ctx, req)
	if err != nil {
		s.errorHandler.Handle(ctx, err)
		return nil, err
	}

	resp, err = s.e(ctx, resp)
	if err != nil {
		s.errorHandler.Handle(ctx, err)
		return nil, err
	}

	for _, f := range s.after {
		ctx = f(ctx, resp)
	}

	resp, err = s.enc(ctx, resp)
	if err != nil {
		s.errorHandler.Handle(ctx, err)
		return nil, err
	}
	return

}
