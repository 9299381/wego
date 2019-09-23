package configs

import (
	"reflect"
)

type Config struct {
}

func (it *Config) GetKey(obj interface{}, key string) string {
	rdata := reflect.ValueOf(obj)
	ret := rdata.Elem().FieldByName(key)
	if ret.IsValid() {
		return ret.Interface().(string)
	} else {
		return ""
	}
}
