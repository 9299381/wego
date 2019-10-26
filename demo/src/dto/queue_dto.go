package dto

import (
	"encoding/json"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/validations"
)

type TaskRequest struct {
	TaskId    string   `json:"task_id" valid:"Required"`
	OrderJson string   `json:"order_json" valid:"Required"`
	OrderList []*Order `json:"-" valid:"Custom(CheckOrder)"`
}

func (s *TaskRequest) CheckOrder(v *validations.Validation) {
	err := json.Unmarshal([]byte(s.OrderJson), &s.OrderList)
	if err != nil {
		_ = v.SetError("order_json", constants.ErrJson)
		return
	}
	for _, order := range s.OrderList {
		err := validations.Valid(order)
		if err != nil {
			_ = v.SetError("list_json", err.Error())
		}
	}
}

type Order struct {
	Id   string `json:"id" valid:"Required" `
	Name string `json:"name" valid:"Required"`
}
