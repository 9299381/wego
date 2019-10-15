package services

import (
	"context"
	"github.com/9299381/wego/contracts"
)

func Pipe() *commonService {
	var s []contracts.IService
	return &commonService{
		serviecs: s,
	}
}

type commonService struct {
	serviecs []contracts.IService
}

func (it *commonService) Middle(s contracts.IService) *commonService {
	it.serviecs = append(it.serviecs, s)
	return it
}
func (it *commonService) Line(ctx contracts.Context) error {
	for _, service := range it.serviecs {
		err := service.Handle(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
func (it *commonService) Parallel(ctx contracts.Context) error {

	type st struct {
		contracts.Context
		err error
	}
	ch := make([]chan st, len(it.serviecs))
	for k, service := range it.serviecs {
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
