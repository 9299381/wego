package ctx

import (
	"fmt"
	"github.com/9299381/wego/contracts"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := &contracts.Context{
		Keys: make(map[string]interface{}),
	}

	ctx.SetValue("a", "1")
	ctx.SetValue("a.b", "2")
	//ctx.SetValue("a", "2")
	fmt.Println(ctx.GetValue("a"))
}
