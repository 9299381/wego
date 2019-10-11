package main

import (
	"fmt"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/clients/mysql"
	"github.com/9299381/wego/demo/src/model"
	"github.com/9299381/wego/tools/idwork"
	"xorm.io/builder"
)

func main() {
	insert()
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
	results := mysql.DB.SQL(sql, args...).Find(&users)
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
	req["id"] = "118916447485100620"
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
	has, _ := mysql.DB.SQL(sql, args...).Get(user)
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
	err := mysql.Fetch(sql, args, &users)
	for _, v := range users {
		fmt.Println(v.Id)
	}
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

func update() {
	user := &model.CommUser{Id: "1189164474851006208"}
	_, _ = mysql.DB.Get(user)
	fmt.Println(user)
	user.UserName = "ccc"
	mysql.Update(user, &model.CommUser{Id: user.Id})
	fmt.Println(user)
}

//insert
func insert() {
	user := &model.CommUser{
		Id:        idwork.ID(),
		UserName:  "go_test",
		Status:    "30",
		LoginName: "aaaaa",
	}
	mysql.Insert(user)

}

//可以创建表
func sync2() {
	_ = mysql.DB.Sync2(new(model.CommUser))
}
