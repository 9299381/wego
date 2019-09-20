package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/configs"
)

type MysqlProvider struct {
}

func (provider *MysqlProvider) Boot() {
	//加载配置文件
	wego.Config("mysql", &configs.MySqlConfig{})

}

func (provider *MysqlProvider) Register() {
	conf := wego.Config("mysql").(*configs.MySqlConfig)
	wego.App.MySql = clients.NewMySql(conf)
}
