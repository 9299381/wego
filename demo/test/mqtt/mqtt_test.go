package mqtt

import (
	"fmt"
	"github.com/9299381/wego/servers/mqtts"
	"testing"
)

func TestMqtt(t *testing.T) {

	m := make(map[string]interface{})
	m["pub"] = "pub"
	m["sub"] = "sub"
	err := mqtts.Publish("sub_test3", m)

	fmt.Println(err)
}
