package json_demo

import (
	"github.com/tidwall/gjson"
	"testing"
)

/**
https://www.jianshu.com/p/623f8ca5ec12
*/
func TestJson(t *testing.T) {
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last").Str
	println(value)
}
