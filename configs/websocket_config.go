package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type WebSocketConfig struct {
	Config
	WebSocketPort string `json:"web_socket_port"`
	Path          string
}

func (it *WebSocketConfig) Load() contracts.Iconfig {

	config := &WebSocketConfig{
		Path:          "/ws",
		WebSocketPort: ":" + wego.Env("SERVER_WEBSOCKET_PORT", "8342"),
	}
	return config
}

func (it *WebSocketConfig) Get(key string) string {
	return it.GetKey(it, key)
}
