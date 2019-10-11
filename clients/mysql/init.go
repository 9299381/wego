package mysql

import (
	"github.com/9299381/wego/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

func init() {
	newMySql()
}

func GetDB() *xorm.Engine {
	return DB
}
func newMySql() {
	conf := (&configs.MySqlConfig{}).Load()
	DB, _ = xorm.NewEngine(
		conf.Driver,
		conf.DataSource,
	)
	DB.SetMaxIdleConns(conf.MaxIdleConns)
	DB.SetMaxOpenConns(conf.MaxOpenConns)
	DB.ShowSQL(conf.ShowSQL)
}
