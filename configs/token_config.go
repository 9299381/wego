package configs

type TokenConfig struct {
	Key string `json:"key"`
	Exp int64  `json:"exp"`
}

func (s *TokenConfig) Load() *TokenConfig {
	config := &TokenConfig{
		Key: EnvString("token.key", "EHKHHP54PXKYTS2E"),
		Exp: int64(EnvInt("token.exp", 2592000)),
	}
	return config
}
