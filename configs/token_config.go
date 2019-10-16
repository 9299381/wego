package configs

import (
	"strconv"
)

type TokenConfig struct {
	Key string `json:"key"`
	Exp int64  `json:"exp"`
}

func (s *TokenConfig) Load() *TokenConfig {
	exp, _ := strconv.Atoi(Env("TOKEN_EXP", "2592000"))
	config := &TokenConfig{
		Key: Env("TOKEN_KEY", "EHKHHP54PXKYTS2E"),
		Exp: int64(exp),
	}
	return config
}
