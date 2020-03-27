package service

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/fsm"
	"github.com/9299381/wego/demo/src/model"
)

type SqlService struct {
}

func (s *SqlService) Handle(ctx contracts.Context) error {

	user := &model.CommUser{Id: "1189164474851006208"}
	_, _ = clients.DB().Get(user)
	ctx.Log.Info(user)

	////初始化状态机
	sm := fsm.NewUserFSM(ctx, user)
	ctx.Log.Info(user.Status)
	//发送状态转换的事件
	if sm.Can("login") {
		err := sm.Event("login")
		if err != nil {
			return err
		}
		user.UserName = "aaaaaaaaa"
		ctx.Log.Info(user.Status)
		_, _ = clients.DB().Update(user, &model.CommUser{Id: user.Id})
	}
	ctx.Set("user", user)

	return nil
}
