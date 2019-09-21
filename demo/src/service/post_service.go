package service

import (
	"github.com/9299381/wego/contracts"
)

type PostService struct {
	next contracts.IService
}

func (it *PostService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *PostService) Handle(ctx contracts.Context) error {

	ctx.Response("request", ctx.Request())

	return it.next.Handle(ctx)
}
