package jcli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/shipengqi/golib/sysutil"
	"github.com/shipengqi/log"
)

func NormalizeCliName(basename string) string {
	if len(basename) == 0 {
		return filepath.Base(os.Args[0])
	}

	if sysutil.IsWindows() {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}

func PrintWorkingDir() {
	wd, _ := os.Getwd()
	log.Infof("%v WorkingDir: %s", progressMessage, wd)
}

type infoLogger struct{}

func (l *infoLogger) Printf(template string, args ...interface{}) {
	log.Infof(template, args...)
}
