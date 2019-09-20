package service

import (
	"github.com/9299381/wego/contracts"
	repository2 "github.com/9299381/wego/example/src/repository"
)

type SqlService struct {
	next contracts.IService
}


func (it *SqlService)Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}
func (it *SqlService)Handle(ctx contracts.Context) error  {
	repo := &repository2.UserRepo{Context: ctx}
	user := repo.FetchId("1189164474851006208")

	ctx.Response("user",user)
	ctx.Response("request",ctx.GetValue("request"))

	return it.next.Handle(ctx)
}
