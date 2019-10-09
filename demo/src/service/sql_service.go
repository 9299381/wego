package service

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/repository"
)

type SqlService struct {
	next contracts.IService
}

func (it *SqlService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}
func (it *SqlService) Handle(ctx contracts.Context) error {
	repo := &repository.UserRepo{Context: ctx}
	user := repo.FetchId("1189164474851006208")
	//初始化状态机
	user.InitFSM()
	ctx.Log.Info(user.Status)
	//发送状态转换的事件
	err := user.FSM.Event("login")
	if err != nil {
		return err
	}
	ctx.Log.Info(user.Status)

	repo.Update(&user)

	ctx.Response("user", user)
	ctx.Response("request", ctx.Request())

	return it.next.Handle(ctx)
}
