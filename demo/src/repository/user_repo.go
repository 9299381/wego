package repository

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/model"
)

type UserRepo struct {
	contracts.Context
}

func (it *UserRepo) FetchId(id string) model.CommUser {
	user := model.CommUser{Id: id}
	has, _ := clients.DB().Get(&user)
	if has {
		return user
	}
	return model.CommUser{}
}
func (it *UserRepo) Update(user *model.CommUser) {
	affected, err := clients.DB().ID(user.Id).Update(user)
	if err != nil {
		it.Log.Info(err)
		return
	}
	it.Log.Info(affected)

}
