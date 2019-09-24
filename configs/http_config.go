package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type HttpConfig struct {
	Config
	HttpHost string `json:"http_host"`
	HttpPort string `json:"http_port"`
}

func (it *HttpConfig) Load() contracts.Iconfig {
	config := &HttpConfig{
		HttpHost: wego.Env("SERVER_HTTP_HOST", "127.0.0.1"),
		HttpPort: wego.Env("SERVER_HTTP_PORT", "8341"),
	}
	return config
}

func (it *HttpConfig) Get(key string) string {
	return it.GetKey(it, key)
}
