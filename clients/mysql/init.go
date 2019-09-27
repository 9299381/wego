package mysql

import (
	"github.com/9299381/wego/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func init() {
	newMySql()
}

func Get() *xorm.Engine {
	return db
}
func newMySql() {
	conf := (&configs.MySqlConfig{}).Load()
	db, _ = xorm.NewEngine(
		conf.Driver,
		conf.DataSource,
	)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.ShowSQL(false)
}
