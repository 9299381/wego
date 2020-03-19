package container

import "github.com/go-macaron/inject"

type Container struct {
	inject.Injector
}

func NewContainer() *Container {
	injector := inject.New()
	return &Container{
		Injector: injector,
	}
}
