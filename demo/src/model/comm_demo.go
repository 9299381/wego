package model

type CommDemoModel struct {
	Id      string `xorm:"pk varchar(21) notnull unique 'id'" json:"id"`
	NumInt1 int    `xorm:"int(11)" json:"num_int1"`
	NumInt2 int    `xorm:"bigint(20)" json:"num_int2"`
	NumBig  int64  `xorm:"bigint(20)" json:"num_big"`
}

func (s *CommDemoModel) TableName() string {
	return "comm_demo"
}
