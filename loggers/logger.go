package loggers

import (
	"github.com/sirupsen/logrus"
)

type logger struct {
	*logrus.Logger
}

func (s logger) Log(keyvals ...interface{}) error {
	s.Info(keyvals)
	return nil
}
