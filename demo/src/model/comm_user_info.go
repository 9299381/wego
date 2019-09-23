package model

type CommUser struct {
	Id        string `json:"id"`
	UserName  string `json:"user_name"`
	LoginName string `json:"login_name"`
}

func (it *CommUser) TableName() string {
	return "comm_user_info"
}
