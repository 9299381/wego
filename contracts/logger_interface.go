package contracts

type ILogger interface {
	Info(keyvals ...interface{})
	Error(keyvals ...interface{})
	Fatal(keyvals ...interface{})
	Trace(keyvals ...interface{})
	Debug(keyvals ...interface{})
	Panic(keyvals ...interface{})

	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}
