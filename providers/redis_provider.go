package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/configs"
)

type RedisProvider struct {
}

func (provider *RedisProvider) Boot() {
	//加载配置文件
	wego.Config("redis", &configs.RedisConfig{})
}

func (provider *RedisProvider) Register() {
	conf := wego.Config("redis").(*configs.RedisConfig)
	wego.App.Redis = clients.NewRedisPool(conf)

}
