package services

import (
	"github.com/9299381/wego/contracts"
)

func Chain(services ...contracts.IService) contracts.IService {
	len := len(services) - 1
	for i := 0; i < len; i++ {
		services[i].Next(services[i+1])
	}
	services[len].Next(&CommService{})
	return services[0]
}
