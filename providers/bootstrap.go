package providers

import (
	"github.com/9299381/wego"
)

type BootStrap struct {
}

func (it *BootStrap) Boot() {
	//环境变量加载
	wego.Provider(&EnvProvider{})
	wego.Provider(&ConsulEnvProvider{})
	//服务注册到配置中心
	wego.Provider(&ConsulRegistyProvider{})

	wego.Provider(&LogProvider{})
	wego.Provider(&MysqlProvider{})
	wego.Provider(&RedisProvider{})
	wego.Provider(&CacheProvider{})

	wego.Provider(&EventProvider{})

}

func (it *BootStrap) Register() {

}
