package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/configs"
)

type AppProvider struct {
}

func (it *AppProvider) Boot() {
	//加载配置文件
	wego.Config("app", &configs.AppConfig{})
}

func (it *AppProvider) Register() {


}
