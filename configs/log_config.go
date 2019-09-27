package configs

type LogConfig struct {
	LogFilePath string `json:"log_file_path"`
	LogFileName string `json:"log_file_name"`
}

func (it *LogConfig) Load() *LogConfig {
	config := &LogConfig{
		LogFilePath: Env("LOG_FILE_PATH", "./logs"),
		LogFileName: Env("LOG_FILE_NAME", "log"),
	}

	return config
}
