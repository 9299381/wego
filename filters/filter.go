package filters

import (
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
)

func New(controller contracts.IController) endpoint.Endpoint {
	return Chain(
		&ResponseEndpoint{},
		&CommEndpoint{Controller: controller},
	)
}

func Auth(controller contracts.IController) endpoint.Endpoint {
	return Chain(
		&ResponseEndpoint{},
		&JwtEndpoint{},
		&CommEndpoint{Controller: controller},
	)
}

func Limit(controller contracts.IController) endpoint.Endpoint {
	return Chain(
		&ResponseEndpoint{},
		&LimitEndpoint{},
		&CommEndpoint{Controller: controller},
	)
}
