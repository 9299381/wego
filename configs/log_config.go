package configs

type LogConfig struct {
	LogFilePath string `json:"log_file_path"`
	LogFileName string `json:"log_file_name"`
}

func (s *LogConfig) Load() *LogConfig {
	config := &LogConfig{
		LogFilePath: EnvString("log.file_path", "./logs"),
		LogFileName: EnvString("log.file_name", "log"),
	}

	return config
}
