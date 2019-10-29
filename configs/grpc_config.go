package configs

type GrpcConfig struct {
	GrpcHost string `json:"grpc_host"`
	GrpcPort string `json:"grpc_port"`
}

func (i *GrpcConfig) Load() *GrpcConfig {
	config := &GrpcConfig{
		GrpcHost: EnvString("server.grpc_host", "127.0.0.1"),
		GrpcPort: EnvString("server.grpc_port", "9341"),
	}
	return config
}
