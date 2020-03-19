package container

import "sync"

var ins *Container
var once sync.Once

func GetIns() *Container {
	once.Do(func() {
		ins = NewContainer()
	})
	return ins
}
