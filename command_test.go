package jcli_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shipengqi/jcli"
)

func TestCommandRun(t *testing.T) {
	t.Run("help message should contain version flag", func(t *testing.T) {
		os.Args = []string{"simplecmd", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.NewCommand("simplecmd", "this is a test command",
			jcli.WithCommandCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.EnableCommandVersion(),
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				log.Infof("command running")
				return nil
			}),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, buf.String(), "[info] command running")
		assert.Contains(t, string(stdout), "--username")
		assert.Contains(t, string(stdout), "--password")
		assert.Contains(t, string(stdout), "--version version[=true]")
		assert.NotContains(t, string(stdout), "-c, --config FILE")
	})

	t.Run("help message should not contain version flag", func(t *testing.T) {
		os.Args = []string{"simplecmd", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.NewCommand("simplecmd", "this is a test command",
			jcli.WithCommandCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				log.Infof("command running")
				return nil
			}),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, buf.String(), "[info] command running")
		assert.Contains(t, string(stdout), "--username")
		assert.Contains(t, string(stdout), "--password")
		assert.NotContains(t, string(stdout), "--version version[=true]")
		assert.NotContains(t, string(stdout), "-c, --config FILE")
	})

	t.Run("with run", func(t *testing.T) {
		os.Args = []string{"simplecmd"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.NewCommand("simplecmd", "this is a test command",
			jcli.WithCommandCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				log.Infof("command running")
				return nil
			}),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.Contains(t, buf.String(), "[info] command running")
		assert.NotContains(t, string(stdout), "--username")
		assert.NotContains(t, string(stdout), "--password")
		assert.NotContains(t, string(stdout), "--version version[=true]")
		assert.NotContains(t, string(stdout), "-c, --config FILE")
	})
}
