package model

import (
	"github.com/9299381/wego/tools"
	"github.com/looplab/fsm"
)

const Begin = "0"
const Reg = "10"
const Login = "20"
const Lock = "30"
const Named = "60"

type CommUser struct {
	Id         string          `xorm:"pk varchar(21) notnull unique 'id'" json:"id"`
	UserName   string          `xorm:"varchar(50) 'user_name'" json:"user_name"`
	LoginName  string          `xorm:"varchar(20)  'login_name'" json:"login_name"`
	Status     string          `xorm:"varchar(2)  'status'" json:"status"`
	CreateTime tools.LocalTime `xorm:"timestamp created 'create_time'" json:"create_time"`
	UpdateTime tools.LocalTime `xorm:"timestamp updated 'update_time'" json:"update_time"`

	FSM *fsm.FSM `xorm:"-" json:"-"`
	//contracts.Context
}

func (it *CommUser) TableName() string {
	return "comm_user_info"
}

func (it *CommUser) InitFSM() *CommUser {
	sm := fsm.NewFSM(
		Begin,
		fsm.Events{
			{Name: "register", Src: []string{Begin}, Dst: Reg},
			{Name: "login", Src: []string{Reg}, Dst: Login},
			{Name: "named", Src: []string{Reg, Login}, Dst: Named},
			{Name: "lock", Src: []string{Reg, Login, Named}, Dst: Lock},
		},
		fsm.Callbacks{
			"after_event": func(e *fsm.Event) {
				//it.Log.Infof("Status to %s is %s\n", it.Status, e.Dst)
				it.Status = e.Dst
			},
		},
	)
	if it.Status != "" {
		sm.SetState(it.Status)
	} else {
		sm.SetState(Begin)
	}
	it.FSM = sm
	return it
}
