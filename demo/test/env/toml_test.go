package env

import (
	"fmt"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/configs"
	"testing"
)

func TestToml(t *testing.T) {
	args.Mode = "dev"
	args.Config = "env"
	id := configs.EnvInt("server_id")
	fmt.Println(id)

}
