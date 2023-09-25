package jcli_test

import (
	"bytes"
	"fmt"
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
	t.Run("help message should contain completion command", func(t *testing.T) {
		os.Args = []string{"simplecmd", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w

		app := jcli.NewCommand("simplecmd", "this is a test command",
			jcli.WithCommandCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.EnableCommandVersion(),
			jcli.EnableCommandCompletion(false),
		)
		app.AddCommands(
			jcli.NewCommand("sub1", "sub1 command description",
				jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
					fmt.Println("sub1 command running")
					return nil
				}),
				jcli.WithCommandDesc("sub1 long desc"),
			),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.Contains(t, string(stdout), "Generate the autocompletion script for the specified shell")
		assert.Contains(t, string(stdout), "sub1 command description")
	})
	t.Run("help message should contain not completion command", func(t *testing.T) {
		os.Args = []string{"simplecmd", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		app := jcli.NewCommand("simplecmd", "this is a test command",
			jcli.WithCommandCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.EnableCommandVersion(),
			jcli.EnableCommandCompletion(true),
		)
		app.AddCommands(
			jcli.NewCommand("sub1", "sub1 command description",
				jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
					fmt.Println("sub1 command running")
					return nil
				}),
				jcli.WithCommandDesc("sub1 long desc"),
			),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, string(stdout), "Generate the autocompletion script for the specified shell")
		assert.Contains(t, string(stdout), "sub1 command description")
	})
	t.Run("help message should not contain completion without EnableCommandCompletion", func(t *testing.T) {
		os.Args = []string{"simplecmd", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		app := jcli.NewCommand("simplecmd", "this is a test command",
			jcli.WithCommandCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.EnableCommandVersion(),
		)
		app.AddCommands(
			jcli.NewCommand("sub1", "sub1 command description",
				jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
					fmt.Println("sub1 command running")
					return nil
				}),
				jcli.WithCommandDesc("sub1 long desc"),
			),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, string(stdout), "Generate the autocompletion script for the specified shell")
		assert.Contains(t, string(stdout), "sub1 command description")
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
