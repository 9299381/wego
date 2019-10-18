package services

import (
	"context"
	"github.com/9299381/wego/contracts"
)

func New(service contracts.IService) *commonService {
	var s []contracts.IService
	s = append(s, service)
	return &commonService{
		services: s,
	}
}

func Pipe() *commonService {
	var s []contracts.IService
	return &commonService{
		services: s,
	}
}

type commonService struct {
	services []contracts.IService
}

func (s *commonService) Middle(services ...contracts.IService) *commonService {
	for _, service := range services {
		s.services = append(s.services, service)
	}
	return s
}

func (s *commonService) Call(ctx contracts.Context) error {
	return s.services[0].Handle(ctx)
}

func (s *commonService) Line(ctx contracts.Context) error {
	for _, service := range s.services {
		err := service.Handle(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *commonService) Parallel(ctx contracts.Context) error {

	type st struct {
		contracts.Context
		err error
	}
	ch := make([]chan st, len(s.services))
	for k, service := range s.services {
		ch[k] = make(chan st)
		cc := contracts.Context{
			Context: context.Background(),
			Log:     ctx.Log,
			Keys:    ctx.Keys,
		}
		go func(cc contracts.Context, s contracts.IService, c chan st) {
			err := s.Handle(cc)
			ret := st{
				Context: cc,
				err:     err,
			}
			c <- ret
		}(cc, service, ch[k])
	}
	m := make(map[string]interface{})
	for _, c := range ch {
		res := <-c
		if res.err != nil {
			return res.err
		}
		for key, value := range res.Keys {
			m[key] = value
		}
	}
	ctx.Keys = m
	return nil
}
