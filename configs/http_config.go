package configs

type HttpConfig struct {
	HttpHost string `json:"http_host"`
	HttpPort string `json:"http_port"`
}

func (it *HttpConfig) Load() *HttpConfig {
	config := &HttpConfig{
		HttpHost: Env("SERVER_HTTP_HOST", "127.0.0.1"),
		HttpPort: Env("SERVER_HTTP_PORT", "8341"),
	}
	return config
}
