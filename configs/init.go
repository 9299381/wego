package configs

import (
	"errors"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/tools/readers"
	"strings"
)

var environment map[string]string

func init() {
	m := strings.Split(args.Config, ",")
	if m[0] == ".env" {
		environment = readFromEnv(m[0])
	} else if m[0] == "registy" {
		environment = readFromConsul(args.Registy)
	} else {
		panic(errors.New("请配置环境变量"))
	}
}
func readFromEnv(filepath string) map[string]string {
	reader := (&readers.IniReader{}).New()
	data := reader.Read(filepath).(map[string]map[string]string)
	ret, _ := data["common"]
	envSection, _ := data[args.Mode]
	for k, v := range envSection {
		ret[k] = v
	}
	return ret
}
func readFromConsul(url string) map[string]string {

	return nil
}
func Env(key string, value ...interface{}) string {
	ret, exist := environment[key]
	if exist && ret != "" {
		return ret
	} else {
		return value[0].(string)
	}
}
