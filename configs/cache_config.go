package configs

type CacheConfig struct {
	Size int `json:"size"`
}

func (s *CacheConfig) Load() *CacheConfig {

	config := &CacheConfig{
		Size: EnvInt("cache.size", 1048576),
	}
	return config
}
