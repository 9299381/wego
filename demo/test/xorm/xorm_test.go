package xorm

import (
	"fmt"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/demo/src/model"
	"github.com/9299381/wego/tools/idwork"
	"testing"
	"xorm.io/builder"
)

func TestFetchListJoin(t *testing.T) {

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

func TestFetchOneJoin(t *testing.T) {
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

func TestFetchOne(t *testing.T) {
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
	has, _ := clients.DB().SQL(sql, args...).Get(user)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(has)
	fmt.Println(user)
}

func TestFetch(t *testing.T) {
	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			OrderBy("id DESC").
			Limit(5, 10).
			ToSQL()

	var users []model.CommUser
	err := clients.DB().SQL(sql, args...).Find(&users)
	for _, v := range users {
		fmt.Println(v.Id)
	}
	fmt.Println(users)
	fmt.Println(err)
}

func TestFage(t *testing.T) {
	page := 1
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

func TestUpdate(t *testing.T) {
	user := &model.CommUser{Id: "1189164474851006208"}
	_, _ = clients.DB().Get(user)
	fmt.Println(user)
	user.UserName = "ccc"
	_, _ = clients.DB().Update(user, &model.CommUser{Id: user.Id})
	fmt.Println(user)
}

//insert
func TestInsert(t *testing.T) {
	user := &model.CommUser{
		Id:        idwork.ID(),
		UserName:  "go_test",
		Status:    "30",
		LoginName: "aaaaa",
	}
	_, _ = clients.DB().Insert(user)

}

//可以创建表
func TestSync2(t *testing.T) {
	err := clients.DB().Sync2(new(model.CommUser))
	if err != nil {
		t.Error(err)
	}
}
