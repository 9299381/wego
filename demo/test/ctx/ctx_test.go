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

	ctx.Set("a", "1")
	ctx.Set("a.b", "2")
	//ctx.Set("a", "2")
	fmt.Println(ctx.Get("a"))
}
