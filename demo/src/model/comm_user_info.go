package model

import (
	"github.com/9299381/wego/tools"
)

type CommUser struct {
	Id         string          `xorm:"pk varchar(21) notnull unique 'id'" json:"id"`
	UserName   string          `xorm:"varchar(50) not null 'user_name'" json:"user_name"`
	LoginName  string          `xorm:"varchar(20)  'login_name'" json:"login_name"`
	Status     string          `xorm:"varchar(2)  'status'" json:"status"`
	CreateTime tools.LocalTime `xorm:"datetime created 'create_time'" json:"create_time"`
	UpdateTime tools.LocalTime `xorm:"datetime updated 'update_time'" json:"update_time"`
}

func (it *CommUser) TableName() string {
	return "comm_user_info"

	//return "comm_user_info_copy1"
}
