package mysql

import (
	"github.com/9299381/wego/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
)

var db *xorm.Engine
var once sync.Once

func GetDB() *xorm.Engine {
	once.Do(func() {
		db = newMySql()
	})
	return db
}
func newMySql() *xorm.Engine {
	conf := (&configs.MySqlConfig{}).Load()
	db, _ = xorm.NewEngine(
		conf.Driver,
		conf.DataSource,
	)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.ShowSQL(conf.ShowSQL)
	return db
}
