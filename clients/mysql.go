package clients

import (
	"github.com/9299381/wego/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func NewMySql(conf *configs.MySqlConfig) (db *xorm.Engine) {
	db, _ = xorm.NewEngine(
		conf.Driver,
		conf.DataSource,
	)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.ShowSQL(false)

	return
}
