package repos

import (
	"encoding/json"
	"github.com/9299381/wego/cache"
)

func GetCache(key string, obj interface{}) error {
	b, err := cache.Get(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return nil
}

func SetCache(key string, obj interface{}, exp int) error {
	return cache.Set(key, obj, exp)
}
