package main

import (
	"fmt"
	"github.com/9299381/wego/contracts"
)

func main() {
	ctx := &contracts.Context{
		Keys: map[string]interface{}{},
	}

	ctx.SetValue("a", "1")
	ctx.SetValue("a.b", "2")
	//ctx.SetValue("a", "2")
	fmt.Println(ctx.GetValue("a"))
}
