package cache

import (
	"encoding/json"
	"github.com/9299381/wego/configs"
	"github.com/coocood/freecache"
	"runtime/debug"
	"strconv"
)

var cache *freecache.Cache

func init() {
	config := (&configs.CacheConfig{}).Load()
	value, err := strconv.Atoi(config.Size)
	if err == nil && value != 0 {
		c := freecache.NewCache(value)
		//根据cache的大小进行设置
		debug.SetGCPercent(20)
		cache = c
	}
}

func Set(key string, value interface{}, exp int) error {
	k := []byte(key)
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = cache.Set(k, v, exp)
	if err != nil {
		return err
	}
	return nil
}
func Get(key string) ([]byte, error) {
	k := []byte(key)
	return cache.Get(k)
}
