package json_demo

import (
	"github.com/tidwall/gjson"
	"testing"
)

func TestJson(t *testing.T) {
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}
