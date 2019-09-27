package loggers

import (
	"github.com/sirupsen/logrus"
)

type logger struct {
	*logrus.Logger
}

func (it logger) Log(keyvals ...interface{}) error {
	it.Info(keyvals)
	return nil
}
