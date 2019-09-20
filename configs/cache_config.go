package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type CacheConfig struct {
	Config
	Size string `json:"size"`
}

func (it *CacheConfig) Load() contracts.Iconfig {

	config := &CacheConfig{
		Size: wego.Env("CACHE_SIZE", "1048576"),
	}
	return config
}

func (it *CacheConfig) Get(key string) string {
	return it.GetKey(it, key)
}
