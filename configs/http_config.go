package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type HttpConfig struct {
	Config
	HttpPort string `json:"http_port"`
}

func (it *HttpConfig) Load() contracts.Iconfig {

	config := &HttpConfig{
		HttpPort: ":" + wego.Env("SERVER_HTTP_PORT", "8341"),
	}
	return config
}

func (it *HttpConfig) Get(key string) string {
	return it.GetKey(it, key)
}
