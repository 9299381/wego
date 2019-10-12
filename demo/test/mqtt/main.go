package main

import (
	"fmt"
	"github.com/9299381/wego/servers/mqtts"
)

func main() {
	m := make(map[string]interface{})
	m["pub"] = "pub"
	m["sub"] = "sub"
	err := mqtts.Publish("sub_test3", m)

	fmt.Println(err)
}
