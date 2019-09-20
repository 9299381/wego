package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"strconv"
)

type TokenConfig struct {
	Config
	Key string `json:"key"`
	Exp int64 `json:"exp"`
}

func (it *TokenConfig) Load() contracts.Iconfig {
	exp, _ := strconv.Atoi(wego.Env("TOKEN_EXP","2592000"))
	config := &TokenConfig{
		Key: wego.Env("TOKEN_KEY","EHKHHP54PXKYTS2E"),
		Exp: int64(exp),

	}
	return config
}

func (it *TokenConfig)Get(key string)  string {
	return it.GetKey(it,key)
}