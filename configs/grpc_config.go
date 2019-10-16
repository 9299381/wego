package configs

type GrpcConfig struct {
	GrpcHost string `json:"grpc_host"`
	GrpcPort string `json:"grpc_port"`
}

func (i *GrpcConfig) Load() *GrpcConfig {
	config := &GrpcConfig{
		GrpcHost: Env("SERVER_GRPC_HOST", "127.0.0.1"),
		GrpcPort: Env("SERVER_GRPC_PORT", "9341"),
	}
	return config
}
