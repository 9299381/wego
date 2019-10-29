package configs

import (
	"fmt"
	"github.com/9299381/wego/args"
)

type MySqlConfig struct {
	Driver       string
	DataSource   string
	MaxIdleConns int
	MaxOpenConns int
	ShowSQL      bool
}

func (s *MySqlConfig) Load() *MySqlConfig {
	driver := EnvString("db.connection", "mysql")
	dataSource := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s"+"?charset=utf8&collation=utf8_general_ci",
		EnvString("db.username", "root"),
		EnvString("db.password", "root"),
		EnvString("db.host", "127.0.0.1"),
		EnvString("db.port", "3306"),
		EnvString("db.database", "default"),
	)
	show := false
	if args.Mode != "prod" {
		show = true
	}
	config := &MySqlConfig{
		Driver:       driver,
		DataSource:   dataSource,
		ShowSQL:      show,
		MaxIdleConns: EnvInt("db.max_idle", 5),
		MaxOpenConns: EnvInt("db.max_open", 50),
	}
	return config
}
