package repository

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/clients/mysql"
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

func (s *UserRepo) First(req map[string]interface{}) (ret *model.CommUser, err error) {
	cond := builder.Eq{}
	for k, v := range req {
		cond[k] = v
	}
	sql, args, _ :=
		builder.
			Select("id,user_name,status,create_time,update_time").
			From("comm_user_info").
			Where(cond).
			ToSQL()

	ret = &model.CommUser{}
	err = mysql.First(sql, args, ret)
	return
}

func (s *UserRepo) Update(user *model.CommUser) {
	affected, err := clients.DB().Update(user, &model.CommUser{Id: user.Id})
	if err != nil {
		s.Log.Info(err)
		return
	}
	s.Log.Info(affected)
}
