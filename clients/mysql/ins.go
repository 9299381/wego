package mysql

import (
	"github.com/9299381/wego/configs"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"xorm.io/xorm"
)

var db *xorm.Engine
var onceMysql sync.Once

func GetDB() *xorm.Engine {
	onceMysql.Do(func() {
		db = newMySql()
	})
	return db
}
func newMySql() *xorm.Engine {
	conf := configs.LoadMySqlConfig()
	engine, _ := xorm.NewEngine(
		conf.Driver,
		conf.DataSource,
	)
	engine.SetMaxIdleConns(conf.MaxIdleConns)
	engine.SetMaxOpenConns(conf.MaxOpenConns)
	engine.ShowSQL(conf.ShowSQL)
	return engine
}
