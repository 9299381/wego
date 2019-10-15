package configs

import (
	"errors"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/tools/readers"
	"os"
	"strings"
)

var environment map[string]string

func init() {
	m := strings.Split(args.Config, ",")
	if m[0] == ".env" {
		path, err := searchEnvFile(m[0])
		if err != nil {
			panic(err)
		} else {
			environment = readFromEnv(path)
		}
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
func searchEnvFile(file string) (string, error) {
	var files []string
	files = append(files, file)
	files = append(files, strings.Join([]string{"../", file}, ""))
	files = append(files, strings.Join([]string{"../../", file}, ""))
	files = append(files, strings.Join([]string{"../env/", file}, ""))
	files = append(files, strings.Join([]string{"./env/", file}, ""))
	for _, v := range files {
		b, _ := pathExists(v)
		if b {
			return v, nil
		}
	}
	return "", errors.New(constants.ErrLoadEnv)
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
