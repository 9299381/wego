package service

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/fsm"
	"github.com/9299381/wego/demo/src/model"
	"github.com/9299381/wego/demo/src/repo"
)

type SqlService struct {
}

func (s *SqlService) Handle(ctx contracts.Context) error {

	st := repo.NewUserRepo(ctx)
	req := make(map[string]interface{})
	req["id"] = "1189164474851006208"
	//req["user_name"] = "aaa"
	user, err := st.GetUser(req)
	ctx.Log.Info(user)
	if err != nil {
		return err
	}

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
		st.Update(user, &model.CommUser{Id: user.Id})
	}
	ctx.Set("user", user)

	return nil
}
