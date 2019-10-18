package repo

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/model"
	"github.com/9299381/wego/repos"
	"xorm.io/builder"
)

type UserRepo struct {
	contracts.Context
	*repos.Repo
}

func NewUserRepo(ctx contracts.Context) *UserRepo {
	return &UserRepo{Context: ctx}
}

func (s *UserRepo) GetUser(req map[string]interface{}) (ret *model.CommUser, err error) {
	cond := builder.Eq{}
	for k, v := range req {
		cond[k] = v
	}
	b := builder.
		Select("id,user_name,status,create_time,update_time").
		From("comm_user_info").
		Where(cond)
	ret = &model.CommUser{}
	err = s.FetchOne(b, ret)
	return
}
