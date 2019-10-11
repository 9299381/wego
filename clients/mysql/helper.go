package mysql

import (
	"errors"
	"github.com/9299381/wego/constants"
)

func First(query interface{}, args []interface{}, bean interface{}) error {
	has, err := DB.SQL(query, args...).Get(bean)
	if err != nil {
		return err
	}
	if !has {
		return errors.New(constants.ErrExsit)
	}
	return nil
}

func Fetch(query interface{}, args []interface{}, bean interface{}) error {
	return DB.SQL(query, args...).Find(bean)
}

func Insert(bean interface{}) bool {
	affected, _ := DB.Insert(bean)
	if affected == 1 {
		return true
	}
	return false
}

func Update(bean interface{}, cond interface{}) bool {
	affected, _ := DB.Update(bean, cond)
	if affected == 1 {
		return true
	}
	return false
}
