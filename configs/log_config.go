package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
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

	return config
}

func (it *LogConfig) Get(key string) string {
	return it.GetKey(it, key)
}
