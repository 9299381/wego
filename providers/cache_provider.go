package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/configs"
	"github.com/coocood/freecache"
	"runtime/debug"
	"strconv"
)

type CacheProvider struct {
}

func (it *CacheProvider) Boot() {
	wego.Config("cache", &configs.CacheConfig{})

}

func (it *CacheProvider) Register() {
	size := wego.Config("cache").Get("Size")
	value, err := strconv.Atoi(size)
	if err == nil && value != 0 {
		cache := freecache.NewCache(value)
		//根据cache的大小进行设置
		debug.SetGCPercent(20)
		wego.App.Cache = cache
	}
}
