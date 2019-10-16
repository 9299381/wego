package configs

type WebSocketConfig struct {
	WebSocketHost string `json:"web_socket_host"`
	WebSocketPort string `json:"web_socket_port"`
	Path          string
}

func (s *WebSocketConfig) Load() *WebSocketConfig {

	config := &WebSocketConfig{
		Path:          "/ws",
		WebSocketPort: Env("SERVER_WEBSOCKET_PORT", "8342"),
		WebSocketHost: Env("SERVER_WEBSOCKET_HOST", "127.0.0.1"),
	}
	return config
}
