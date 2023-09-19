package jcli_test

import (
	"fmt"
	"io"

	"github.com/shipengqi/jcli"
)

var (
	debugLevel = []byte{'d', 'e', 'b', 'u', 'g'}
	infoLevel  = []byte{'i', 'n', 'f', 'o'}
	warnLevel  = []byte{'w', 'a', 'r', 'n'}
	errorLevel = []byte{'e', 'r', 'r', 'o', 'r'}
	fatalLevel = []byte{'f', 'a', 't', 'a', 'l'}
)

const levelTmpl = "[%s] "

type testLogger struct {
	wr io.Writer
}

func newTestLogger(out io.Writer) jcli.Logger {
	return &testLogger{wr: out}
}

func (l *testLogger) Debugf(template string, args ...interface{}) {
	l.write(debugLevel, template, args...)
}

func (l *testLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.write(debugLevel, msg, keysAndValues...)
}

func (l *testLogger) Infof(template string, args ...interface{}) {
	l.write(infoLevel, template, args...)
}

func (l *testLogger) Info(msg string, keysAndValues ...interface{}) {
	l.write(infoLevel, msg, keysAndValues...)
}

func (l *testLogger) Warnf(template string, args ...interface{}) {
	l.write(warnLevel, template, args...)
}

func (l *testLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.write(warnLevel, msg, keysAndValues...)
}

func (l *testLogger) Errorf(template string, args ...interface{}) {
	l.write(errorLevel, template, args...)
}

func (l *testLogger) Error(msg string, keysAndValues ...interface{}) {
	l.write(errorLevel, msg, keysAndValues...)
}

func (l *testLogger) Fatalf(template string, args ...interface{}) {
	l.write(fatalLevel, template, args...)
}

func (l *testLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.write(fatalLevel, msg, keysAndValues...)
}

func (l *testLogger) write(level []byte, template string, args ...interface{}) {
	_, _ = fmt.Fprintf(l.wr, levelTmpl, level)
	_, _ = fmt.Fprintf(l.wr, template, args...)
	_, _ = fmt.Fprint(l.wr, "\n")
}
