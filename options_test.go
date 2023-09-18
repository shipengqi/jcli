package jcli_test

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shipengqi/jcli"
)

func TestAppOptions(t *testing.T) {
	t.Run("with basename option", func(t *testing.T) {
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("testapp",
			jcli.WithBaseName("testappbasename"),
			jcli.WithLogger(log),
		)
		app.Run()
		assert.Contains(t, buf.String(), "testappbasename")
		assert.Contains(t, buf.String(), "WorkingDir:")
		assert.Contains(t, buf.String(), "Starting testapp")

	})

	t.Run("with silence", func(t *testing.T) {
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("testapp",
			jcli.WithBaseName("testappbasename"),
			jcli.WithSilence(),
			jcli.WithLogger(log),
		)
		app.Run()
		assert.Contains(t, buf.String(), "testappbasename")
		assert.NotContains(t, buf.String(), "WorkingDir:")
		assert.NotContains(t, buf.String(), "Starting testapp")
	})
}

func TestCommandOptions(t *testing.T) {
	t.Run("with desc and example", func(t *testing.T) {
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w

		cmd := jcli.NewCommand(
			"testcmd",
			"this is a test command",
			jcli.WithCommandDesc("this is a test description for command"),
			jcli.WithCommandExamples("this is a test example for command"),
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				return cmd.Help()
			}))
		cmd.Run()
		expected := []string{
			"this is a test description for command",
			"",
			"Usage:",
			"  testcmd [flags]",
			"",
			"Examples:",
			"  this is a test example for command",
		}
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		reader := bytes.NewReader(stdout)
		scanner := bufio.NewScanner(reader)
		for _, v := range expected {
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			assert.Equal(t, v, line)
		}
	})

	t.Run("with alias", func(t *testing.T) {
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w

		cmd := jcli.NewCommand(
			"testcmd",
			"this is a test command",
			jcli.WithCommandDesc("this is a test description for command"),
			jcli.WithCommandExamples("this is a test example for command"),
			jcli.WithCommandAliases("alias1", "alias2"),
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				return cmd.Help()
			}))
		cmd.Run()
		expected := []string{
			"this is a test description for command",
			"",
			"Usage:",
			"  testcmd [flags]",
			"",
			"Aliases:",
			"  testcmd, alias1, alias2",
			"",
			"Examples:",
			"  this is a test example for command",
		}
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		reader := bytes.NewReader(stdout)
		scanner := bufio.NewScanner(reader)
		for _, v := range expected {
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			assert.Equal(t, v, line)
		}
	})
}
