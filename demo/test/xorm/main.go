package main

import (
	"fmt"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/demo/src/model"
	"xorm.io/builder"
)

func main() {
	page()
}

func fetchListJoin() {

	sql, args, _ :=
		builder.
			Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
			From("comm_user_info as t1").
			LeftJoin("user_bank as t2", "t1.user_id=t2.user_id").
			Limit(10, 0).
			ToSQL()

	type Detail struct {
		Id     string
		CardNo string
	}
	type UserDetail struct {
		model.CommUser `xorm:"extends"`
		Detail         `xorm:"extends"`
	}
	var users []UserDetail
	results := clients.DB().SQL(sql, args...).Find(&users)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(results)
	fmt.Println(users)
}

func fetchOneJoin() {
	sql, args, _ :=
		builder.
			Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
			From("comm_user_info as t1").
			LeftJoin("user_bank as t2", "t1.user_id=t2.user_id").
			Where(builder.Eq{"t1.id": "1189164474851006208"}).
			And(builder.Eq{"t1.user_name": "aaaaaaaaa"}).
			ToSQL()

	type Detail struct {
		Id     string
		CardNo string
	}
	type UserDetail struct {
		model.CommUser `xorm:"extends"`
		Detail         `xorm:"extends"`
	}
	user := &UserDetail{}
	results, _ := clients.DB().SQL(sql, args...).Get(user)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(results)
	fmt.Println(user.CardNo)
}

func fetchOne() {
	req := make(map[string]interface{})
	req["id"] = "1189164474851006208"
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

	user := &model.CommUser{}
	has, _ := clients.DB().SQL(sql, args...).Get(user)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(has)
	fmt.Println(user)
}

func fetch() {
	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			OrderBy("id DESC").
			Limit(5, 10).
			ToSQL()

	var users []model.CommUser
	err := clients.DB().
		SQL(sql, args...).
		Find(&users)

	fmt.Println(users)
	fmt.Println(err)
}

func page() {
	page := 0 //页数
	if page-1 <= 0 {
		page = 1
	}
	pageSize := 10 //页面大小

	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			OrderBy("id DESC").
			Limit(pageSize*(page), (page-1)*pageSize).
			ToSQL()

	var users []model.CommUser
	err := clients.DB().
		SQL(sql, args...).
		Find(&users)

	fmt.Println(users)
	fmt.Println(err)
}
