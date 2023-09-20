package jcli_test

import (
	"bytes"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shipengqi/jcli"
)

func TestAppSignalReceiver(t *testing.T) {
	t.Run("with default signals", func(t *testing.T) {
		os.Args = []string{"testApp"}
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		var pid int
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithLogger(log),
			jcli.DisableConfig(),
			jcli.DisableVersion(),
			jcli.WithRunFunc(func() error {
				log.Infof("application running")
				time.Sleep(time.Second * 2)
				return nil
			}),
			jcli.WithOnSignalReceived(func(signal os.Signal) {
				log.Infof("signal: %s", signal.String())
			}),
		)
		go func() {
			pid = syscall.Getpid()
			app.Run()
		}()
		time.Sleep(time.Millisecond * 500)
		go func() {
			syscall.Kill(pid, syscall.SIGINT)
		}()
		time.Sleep(time.Second * 5)
		assert.Contains(t, buf.String(), "[info] application running")
		assert.Contains(t, buf.String(), "signal: interrupt")
	})
	t.Run("with custom signals", func(t *testing.T) {
		os.Args = []string{"testApp"}
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		var pid int
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithLogger(log),
			jcli.DisableConfig(),
			jcli.DisableVersion(),
			jcli.WithRunFunc(func() error {
				log.Infof("application running")
				time.Sleep(time.Second * 2)
				return nil
			}),
			jcli.WithOnSignalReceived(func(signal os.Signal) {
				log.Infof("signal: %s", signal.String())
			}, syscall.SIGUSR1),
		)
		go func() {
			pid = syscall.Getpid()
			app.Run()
		}()
		time.Sleep(time.Millisecond * 500)
		go func() {
			syscall.Kill(pid, syscall.SIGUSR1)
		}()
		time.Sleep(time.Second * 5)
		assert.Contains(t, buf.String(), "[info] application running")
		assert.Contains(t, buf.String(), "signal: user defined signal 1")
	})
}
