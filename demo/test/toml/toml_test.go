package toml

import (
	"fmt"
	"github.com/9299381/wego"
	"testing"
)

func TestToml(t *testing.T) {
	// 加载过程
	wego.Toml("user", "user")
	// 读取过程
	level := wego.Toml("user").GetString("type.register")
	fmt.Println(level)

}
