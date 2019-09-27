package repository

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/model"
)

type UserRepo struct {
	contracts.Context
}

func (it UserRepo) FetchId(id string) model.CommUser {

	user := model.CommUser{Id: id}
	has, _ := clients.DB().Get(&user)
	it.Log.Info(user)
	if has {
		return user
	}
	return model.CommUser{}
}
