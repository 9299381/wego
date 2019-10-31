package repos

import (
	"errors"
	"github.com/9299381/wego/clients/mysql"
	"github.com/9299381/wego/constants"
	"github.com/go-xorm/xorm"
	"xorm.io/builder"
)

func DB() *xorm.Engine {
	return mysql.GetDB()
}

func First(obj interface{}) error {
	has, err := mysql.GetDB().Get(obj)
	if err != nil {
		return err
	}
	if !has {
		return errors.New(constants.ErrNotExist)
	}
	return nil
}
func Find(beans interface{}, condiBeans ...interface{}) error {
	return mysql.GetDB().Find(beans, condiBeans)
}

func Insert(bean interface{}) bool {
	affected, _ := mysql.GetDB().Insert(bean)
	if affected == 1 {
		return true
	}
	return false
}

func Update(bean interface{}, cond interface{}) bool {
	affected, _ := mysql.GetDB().Update(bean, cond)
	if affected == 1 {
		return true
	}
	return false
}
func Exist(obj interface{}) bool {
	b, err := mysql.GetDB().Exist(obj)
	if err != nil {
		return false
	}
	return b
}

func FetchOne(b *builder.Builder, bean interface{}) error {
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	has, err := mysql.GetDB().SQL(query, args...).Get(bean)
	if err != nil {
		return err
	}
	if !has {
		return errors.New(constants.ErrNotExist)
	}
	return nil
}

func Fetch(b *builder.Builder, bean interface{}) error {
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	return mysql.GetDB().SQL(query, args...).Find(bean)
}
