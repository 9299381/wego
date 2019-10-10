package repository

import (
	"errors"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/model"
	"xorm.io/builder"
)

type UserRepo struct {
	contracts.Context
}

func NewUserRepo(ctx contracts.Context) *UserRepo {
	return &UserRepo{Context: ctx}
}

func (it *UserRepo) Get(req map[string]interface{}) (ret *model.CommUser, err error) {
	cond := builder.Eq{}
	for k, v := range req {
		cond[k] = v
	}
	sql, args, _ := builder.MySQL().
		Select("id,user_name,status,create_time,update_time").
		From("comm_user_info").
		Where(cond).
		ToSQL()

	ret = &model.CommUser{}
	has, err := clients.DB().SQL(sql, args...).Get(ret)
	if err != nil {
		return
	}
	if has {
		return
	} else {
		err = errors.New("用户不存在")
		return
	}
}

func (it *UserRepo) Update(user *model.CommUser) {
	affected, err := clients.DB().ID(user.Id).Update(user)
	if err != nil {
		it.Log.Info(err)
		return
	}
	it.Log.Info(affected)
}
