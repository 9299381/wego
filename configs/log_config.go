package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/errors"
	"os"
)

type LogConfig struct {
	Config
	LogFilePath string `json:"log_file_path"`
	LogFileName string `json:"log_file_name"`
}

func (it *LogConfig) Load() contracts.Iconfig {

	config := &LogConfig{
		LogFilePath: wego.Env("LOG_FILE_PATH", "/logs"),
		LogFileName: wego.Env("LOG_FILE_NAME", "log"),
	}
	exist, err := it.pathExists(config.LogFilePath)
	if err != nil {
		panic(errors.New("9999", "日志目录配置有问题"))
	}
	if !exist {
		err := os.Mkdir(config.LogFilePath, os.ModePerm)
		if err != nil {
			panic(errors.New("9999", "创建日志目录失败"))
		}
	}

	return config
}

func (it *LogConfig) Get(key string) string {
	return it.GetKey(it, key)
}

func (it *LogConfig) pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
