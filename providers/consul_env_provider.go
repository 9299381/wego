package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"strings"
)

type ConsulEnvProvider struct {
}

func (it *ConsulEnvProvider) Boot() {
	if strings.Contains(args.Config, "consul:") {
		wego.App.Env = it.ReadConsul()
	}
}

func (it *ConsulEnvProvider) Register() {

}

func (it *ConsulEnvProvider) ReadConsul() map[string]string {
	//todo 从consul服务器读取env配置
	return nil
}
