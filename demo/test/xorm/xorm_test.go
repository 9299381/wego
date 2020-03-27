package xorm

import (
	"fmt"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/demo/src/model"
	"github.com/9299381/wego/tools/idwork"
	"testing"
	"xorm.io/builder"
)

type Detail struct {
	Id     string
	CardNo string
}
type UserDetail struct {
	model.CommUser `xorm:"extends"`
	Detail         `xorm:"extends"`
}

func TestListJoin(t *testing.T) {

	var users []UserDetail
	results := clients.DB().
		Table("comm_user_info").
		Alias("t1").
		Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
		Join("LEFT", "user_bank as t2", "t1.user_id=t2.user_id").
		Limit(10, 0).
		Find(&users)
	fmt.Println(results)
	fmt.Println(users)
	for key, value := range users {
		fmt.Println(key)
		fmt.Println(value.UserName)
	}
}

func TestOneJoin(t *testing.T) {
	user := &UserDetail{}
	result, err := clients.DB().
		Table("comm_user_info").
		Alias("t1").
		Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
		Join("LEFT", "user_bank as t2", "t1.user_id=t2.user_id").
		Where("t1.id = ?", "1189164474851006208").
		And("t1.user_name = ?", "ccc").
		Limit(1).
		Get(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(user.CardNo)
}

func TestOne(t *testing.T) {
	user := &model.CommUser{}
	result, _ := clients.DB().
		//Table("comm_user_info").
		Select("id,user_name,status,create_time,update_time").
		Where("id =?", "1189164474851006208").
		Limit(1).
		Get(user)
	fmt.Println(result)
	fmt.Println(user)
}
func TestGet(t *testing.T) {
	user := &model.CommUser{Id: "1189164474851006208"}
	result, _ := clients.DB().Get(user)
	fmt.Println(result)
	fmt.Println(user)
}

func TestList(t *testing.T) {
	var users []model.CommUser
	results := clients.DB().
		Table("comm_user_info").
		Select("id,user_name,status,create_time,update_time").
		OrderBy("id DESC").
		Limit(5, 10).
		Find(&users)
	fmt.Println(results)
	fmt.Println(users)
	for key, value := range users {
		fmt.Println(key)
		fmt.Println(value.UserName)
	}
}

func TestPage(t *testing.T) {
	var users []model.CommUser
	page := 1
	pageSize := 10 //页面大小
	results := clients.DB().
		Table("comm_user_info").
		Select("id,user_name,status,create_time,update_time").
		Where("status = ?", "20").
		OrderBy("id DESC").
		Limit(pageSize*(page), (page-1)*pageSize).
		Find(&users)
	fmt.Println(results)
	fmt.Println(users)
	for key, value := range users {
		fmt.Println(key)
		fmt.Println(value.UserName)
	}
}

func TestUpdateOne(t *testing.T) {
	user := &model.CommUser{Id: "1306582895206334464"}
	//_, _ = clients.DB().Get(user)
	//fmt.Println(user)
	user.UserName = "ccc"
	_, _ = clients.DB().Update(user, &model.CommUser{Id: user.Id})
	fmt.Println(user)
}

func TestUpdateTwo(t *testing.T) {
	user := &model.CommUser{Id: "1306582895206334464"}
	//_, _ = clients.DB().Get(user)
	//fmt.Println(user)
	user.UserName = "ccc"
	_, _ = clients.DB().
		ID(user.Id).
		Cols("user_name").
		Update(user)
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

/// builder sql 方式

func TestBuilderFetchListJoin(t *testing.T) {

	sql, args, _ :=
		builder.
			Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
			From("comm_user_info as t1").
			LeftJoin("user_bank as t2", "t1.user_id=t2.user_id").
			Limit(10, 0).
			ToSQL()
	var users []UserDetail
	results := clients.DB().SQL(sql, args...).Find(&users)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(results)
	fmt.Println(users)
}

func TestBuilderFetchOneJoin(t *testing.T) {
	sql, args, _ :=
		builder.
			Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
			From("comm_user_info as t1").
			LeftJoin("user_bank as t2", "t1.user_id=t2.user_id").
			Where(builder.Eq{"t1.id": "1189164474851006208"}).
			And(builder.Eq{"t1.user_name": "aaaaaaaaa"}).
			ToSQL()

	user := &UserDetail{}
	results, _ := clients.DB().SQL(sql, args...).Get(user)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(results)
	fmt.Println(user.CardNo)
}
func TestBuilderFetchOne(t *testing.T) {
	req := make(map[string]interface{})
	req["id"] = "1306582895206334464"
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

func TestBuilderFetch(t *testing.T) {
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
func TestBuilderFage(t *testing.T) {
	page := 1
	pageSize := 10 //页面大小

	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			Where(builder.Eq{"status": "20"}).
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

func TestBuilderPageList(t *testing.T) {
	page := 1
	pageSize := 10 //页面大小

	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			Where(builder.Eq{"status": "20"}).
			OrderBy("id DESC").
			ToSQL()
	var users []model.CommUser
	err := clients.DB().
		SQL(sql, args...).
		Limit(pageSize*(page), (page-1)*pageSize).
		Find(&users)

	fmt.Println(users)
	fmt.Println(err)
}
