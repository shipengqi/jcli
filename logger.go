package jcli

type Logger interface {
	Debugf(template string, args ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Infof(template string, args ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warnf(template string, args ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Errorf(template string, args ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatalf(template string, args ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}
