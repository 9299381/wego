package configs

import (
	"fmt"
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/contracts"
	"strconv"
)

type MySqlConfig struct {
	Config
	Driver       string
	DataSource   string
	MaxIdleConns int
	MaxOpenConns int
	ShowSQL      bool
}

func (it *MySqlConfig) Load() contracts.Iconfig {
	driver := wego.Env("DB_CONNECTION", "mysql")
	dataSource := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s"+"?charset=utf8&collation=utf8_general_ci",
		wego.Env("DB_USERNAME", "root"),
		wego.Env("DB_PASSWORD", "root"),
		wego.Env("DB_HOST", "127.0.0.1"),
		wego.Env("DB_PORT", "3306"),
		wego.Env("DB_DATABASE", "default"),
	)

	show := false
	if args.Mode != "prod" {
		show = true
	}

	maxIdel, _ := strconv.Atoi(wego.Env("DB_MAX_IDLE", "5"))
	maxOpen, _ := strconv.Atoi(wego.Env("DB_MAX_OPEN", "50"))

	config := &MySqlConfig{
		Driver:       driver,
		DataSource:   dataSource,
		ShowSQL:      show,
		MaxIdleConns: maxIdel,
		MaxOpenConns: maxOpen,
	}
	return config
}

func (it *MySqlConfig) Get(key string) string {
	return it.GetKey(it, key)
}
