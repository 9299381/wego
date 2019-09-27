package loggers

import (
	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = newLogrus()
}

// log.logger 接口的一种实现,用以注入 go-kit 的服务注册
func NewKitLog() log.Logger {
	return logger{
		Logger: Log,
	}
}
