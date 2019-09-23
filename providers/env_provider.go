package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/readers"
	"strings"
)

type EnvProvider struct {
}

func (it *EnvProvider) Boot() {
	if strings.Contains(args.Config, "env") {
		wego.App.Env = it.ReadEnv(&readers.IniReader{})
	}
}

func (it *EnvProvider) Register() {

}

func (it *EnvProvider) ReadEnv(reader contracts.IReader) map[string]interface{} {
	data := reader.Read(args.Config).(map[string]map[string]interface{})
	ret, _ := data["common"]

	envSection, _ := data[args.Mode]
	for k, v := range envSection {
		ret[k] = v
	}
	return ret
}
