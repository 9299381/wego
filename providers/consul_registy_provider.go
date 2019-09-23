package providers

import (
	"github.com/9299381/wego/args"
)

type ConsulRegistyProvider struct {
}

func (it *ConsulRegistyProvider) Boot() {

}

func (it *ConsulRegistyProvider) Register() {
	if args.Registy != "" {
		//把自己作为服务注册到注册中心
		//需要判断启动服务是 http,grpc
		// serviceName = args.Name + "_" + http
		// serviceId = 随机 wego.ID
	}
}
