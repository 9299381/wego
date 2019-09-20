package providers

import "github.com/9299381/wego"

type BootStrap struct {

}

func (it *BootStrap) Boot() {
	wego.Provider(&AppProvider{})
	wego.Provider(&LogProvider{})
	wego.Provider(&MysqlProvider{})
	wego.Provider(&RedisProvider{})
	wego.Provider(&CacheProvider{})
}

func (it *BootStrap) Register() {


}
