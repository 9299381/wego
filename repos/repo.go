package repos

import (
	"encoding/json"
	"errors"
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/clients/mysql"
	"github.com/9299381/wego/constants"
	"github.com/go-xorm/xorm"
	"xorm.io/builder"
)

func New() *Repo {
	return &Repo{}
}

type Repo struct {
}

func (s *Repo) DB() *xorm.Engine {
	return mysql.GetDB()
}

func (s *Repo) First(obj interface{}) (bool, error) {
	return mysql.GetDB().Get(obj)
}
func (s *Repo) Find(beans interface{}, condiBeans ...interface{}) error {
	return mysql.GetDB().Find(beans, condiBeans)
}

func (s *Repo) Insert(bean interface{}) bool {
	affected, _ := mysql.GetDB().Insert(bean)
	if affected == 1 {
		return true
	}
	return false
}

func (s *Repo) Update(bean interface{}, cond interface{}) bool {
	affected, _ := mysql.GetDB().Update(bean, cond)
	if affected == 1 {
		return true
	}
	return false
}
func (s *Repo) Exist(obj interface{}) (bool, error) {
	return mysql.GetDB().Exist(obj)
}

func (s *Repo) FetchOne(b *builder.Builder, bean interface{}) error {
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	has, err := mysql.GetDB().SQL(query, args...).Get(bean)
	if err != nil {
		return err
	}
	if !has {
		return errors.New(constants.ErrExsit)
	}
	return nil
}

func (s *Repo) Fetch(b *builder.Builder, bean interface{}) error {
	query, args, err := b.ToSQL()
	if err != nil {
		return err
	}
	return mysql.GetDB().SQL(query, args...).Find(bean)
}

//////////////

func (s *Repo) GetCache(key string, obj interface{}) error {
	b, err := cache.Get(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repo) SetCache(key string, obj interface{}, exp int) error {
	return cache.Set(key, obj, exp)
}
