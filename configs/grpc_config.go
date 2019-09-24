package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type GrpcConfig struct {
	Config
	GrpcHost string `json:"grpc_host"`
	GrpcPort string `json:"grpc_port"`
}

func (it *GrpcConfig) Load() contracts.Iconfig {
	config := &GrpcConfig{
		GrpcHost: wego.Env("SERVER_GRPC_HOST", "127.0.0.1"),
		GrpcPort: wego.Env("SERVER_GRPC_PORT", "9341"),
	}
	return config
}

func (it *GrpcConfig) Get(key string) string {
	return it.GetKey(it, key)
}
