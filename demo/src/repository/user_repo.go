package repository

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/model"
)

type UserRepo struct {
	contracts.Context
}

func (it UserRepo) FetchId(id string) model.CommUser {

	user := model.CommUser{Id: id, UserName: "系统二级1111"}
	has, _ := wego.DB().Get(&user)
	it.Log.Info("user_repo  fetch id")
	it.Log.Info(user)
	if has {
		return user
	}
	return model.CommUser{}
}
