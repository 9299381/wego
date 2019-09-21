package service

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

type AuthService struct {
	next contracts.IService
}

func (it *AuthService) Next(srv contracts.IService) contracts.IService {
	//这里可以做些事情
	it.next = srv
	return it
}

func (it *AuthService) Handle(ctx contracts.Context) error {

	fmt.Println(ctx.GetValue("request.claim.Id"))
	fmt.Println(ctx.GetValue("request.claim.Name"))

	return it.next.Handle(ctx)
}
