package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/configs"
	"strings"
)

type ConsulRegistyProvider struct {
}

func (it *ConsulRegistyProvider) Boot() {

}

func (it *ConsulRegistyProvider) Register() {
	if args.Registy != "" {
		if strings.Contains(args.Server, "http") || strings.Contains(args.Server, "gateway") {
			httpConfig := (&configs.HttpConfig{}).Load()
			wego.App.Consul["http"] = clients.NewConsulHttpRegister(
				args.Name,
				httpConfig.HttpHost,
				httpConfig.HttpPort,
			)
		}
		if strings.Contains(args.Server, "grpc") {
			grpcConfig := (&configs.GrpcConfig{}).Load()
			wego.App.Consul["grpc"] = clients.NewConsulGrpcRegister(
				args.Name,
				grpcConfig.GrpcHost,
				grpcConfig.GrpcPort,
			)
		}
	}
}
