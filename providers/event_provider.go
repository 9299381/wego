package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type EventProvider struct {
}

func (it *EventProvider) Boot() {
	//加载配置文件,事件侦听频率
	//wego.Config("cache", &configs.CacheConfig{})

}

func (it *EventProvider) Register() {
	//内置加载事件服务,无需路由,直接调用 filter handler
	wego.Router("event", servers.NewEventCommServer(3))
}
