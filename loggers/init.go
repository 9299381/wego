package loggers

import (
	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
	"sync"
)

var ins *logrus.Logger
var once sync.Once

func GetLog() *logrus.Logger {
	once.Do(func() {
		ins = newLogrus()
	})
	return ins
}

// log.logger 接口的一种实现,用以注入 go-kit 的服务注册
func NewKitLog() log.Logger {
	return logger{
		Logger: GetLog(),
	}
}
