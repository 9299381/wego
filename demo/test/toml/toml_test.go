package toml

import (
	"fmt"
	"github.com/9299381/wego"
	"testing"
)

func TestToml(t *testing.T) {
	wego.Toml("user", "user")
	level := wego.Toml("user").GetString("type.register")
	fmt.Println(level)

}
