package fsm

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/model"
	"github.com/looplab/fsm"
)

const Begin = "0"
const Reg = "10"
const Login = "20"
const Lock = "30"
const Named = "60"

type UserFSM struct {
	contracts.Context
	*fsm.FSM
	User *model.CommUser
}

func NewUserFSM(ctx contracts.Context, user *model.CommUser) *UserFSM {
	it := &UserFSM{
		Context: ctx,
		User:    user,
	}
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
				it.Log.Infof("fsm event:%s change status to %s", e.Event, e.Dst)
				it.User.Status = e.Dst
			},
		},
	)
	if it.User.Status != "" {
		sm.SetState(it.User.Status)
	} else {
		sm.SetState(Begin)
	}

	it.FSM = sm
	return it
}
