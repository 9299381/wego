package configs

import (
	"fmt"
	"github.com/9299381/wego/args"
	"strconv"
)

type MySqlConfig struct {
	Driver       string
	DataSource   string
	MaxIdleConns int
	MaxOpenConns int
	ShowSQL      bool
}

func (s *MySqlConfig) Load() *MySqlConfig {
	driver := Env("DB_CONNECTION", "mysql")
	dataSource := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s"+"?charset=utf8&collation=utf8_general_ci",
		Env("DB_USERNAME", "root"),
		Env("DB_PASSWORD", "root"),
		Env("DB_HOST", "127.0.0.1"),
		Env("DB_PORT", "3306"),
		Env("DB_DATABASE", "default"),
	)

	show := false
	if args.Mode != "prod" {
		show = true
	}

	maxIdel, _ := strconv.Atoi(Env("DB_MAX_IDLE", "5"))
	maxOpen, _ := strconv.Atoi(Env("DB_MAX_OPEN", "50"))

	config := &MySqlConfig{
		Driver:       driver,
		DataSource:   dataSource,
		ShowSQL:      show,
		MaxIdleConns: maxIdel,
		MaxOpenConns: maxOpen,
	}
	return config
}
