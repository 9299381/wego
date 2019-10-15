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
	engine, _ := xorm.NewEngine(
		conf.Driver,
		conf.DataSource,
	)
	engine.SetMaxIdleConns(conf.MaxIdleConns)
	engine.SetMaxOpenConns(conf.MaxOpenConns)
	engine.ShowSQL(conf.ShowSQL)
	return engine
}
