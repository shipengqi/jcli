package jcli_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/shipengqi/jcli"
)

type fakeCliOptions struct {
	Username string
	Password string
}

func (o *fakeCliOptions) Flags() (fss cliflag.NamedFlagSets) {
	fakes := fss.FlagSet("fake")
	fakes.StringVar(&o.Username, "username", o.Username, "fake username.")
	fakes.StringVar(&o.Password, "password", o.Password, "fake password.")

	return fss
}

func (o *fakeCliOptions) Validate() []error {
	return nil
}

func TestAppRun(t *testing.T) {
	t.Run("help message should contain version and config flag", func(t *testing.T) {
		os.Args = []string{"testApp", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithLogger(log),
			jcli.WithRunFunc(func() error {
				log.Infof("application running")
				return nil
			}))
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, buf.String(), "[info] application running")
		assert.Contains(t, string(stdout), "--username")
		assert.Contains(t, string(stdout), "--password")
		assert.Contains(t, string(stdout), "--version version[=true]")
		assert.Contains(t, string(stdout), "-c, --config FILE")
	})
	t.Run("help message should contain examples", func(t *testing.T) {
		os.Args = []string{"testApp", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithExamples("This is a example for testing"),
			jcli.WithLogger(log),
			jcli.WithRunFunc(func() error {
				log.Infof("application running")
				return nil
			}))
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, buf.String(), "[info] application running")
		assert.Contains(t, string(stdout), "This is a example for testing")
		assert.Contains(t, string(stdout), "--username")
		assert.Contains(t, string(stdout), "--password")
		assert.Contains(t, string(stdout), "--version version[=true]")
		assert.Contains(t, string(stdout), "-c, --config FILE")
	})
	t.Run("help message should contain aliases", func(t *testing.T) {
		os.Args = []string{"testApp", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithExamples("This is a example for testing"),
			jcli.WithAliases("alias1", "alias2"),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.Contains(t, string(stdout), "This is a example for testing")
		assert.Contains(t, string(stdout), "alias1, alias2")
		assert.Contains(t, string(stdout), "--username")
		assert.Contains(t, string(stdout), "--password")
		assert.Contains(t, string(stdout), "--version version[=true]")
		assert.Contains(t, string(stdout), "-c, --config FILE")
	})
	t.Run("help message should contain description", func(t *testing.T) {
		os.Args = []string{"testApp", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w

		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithExamples("This is a example for testing"),
			jcli.WithDesc("This is a description for testing"))
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.Contains(t, string(stdout), "This is a example for testing")
		assert.Contains(t, string(stdout), "This is a description for testing")
		assert.Contains(t, string(stdout), "--username")
		assert.Contains(t, string(stdout), "--password")
		assert.Contains(t, string(stdout), "--version version[=true]")
		assert.Contains(t, string(stdout), "-c, --config FILE")
	})
	t.Run("help message should not contain version flag", func(t *testing.T) {
		os.Args = []string{"testApp", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithLogger(log),
			jcli.DisableVersion(),
			jcli.WithRunFunc(func() error {
				log.Infof("application running")
				return nil
			}))
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, buf.String(), "[info] application running")
		assert.NotContains(t, string(stdout), "--version version[=true]")
		assert.Contains(t, string(stdout), "-c, --config FILE")
	})
	t.Run("help message should not contain config flag", func(t *testing.T) {
		os.Args = []string{"testApp", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithLogger(log),
			jcli.DisableConfig(),
			jcli.DisableVersion(),
			jcli.WithRunFunc(func() error {
				log.Infof("application running")
				return nil
			}))
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, buf.String(), "[info] application running")
		assert.NotContains(t, string(stdout), "--version version[=true]")
		assert.NotContains(t, string(stdout), "-c, --config FILE")
	})
	t.Run("with run", func(t *testing.T) {
		os.Args = []string{"testApp"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithLogger(log),
			jcli.DisableConfig(),
			jcli.DisableVersion(),
			jcli.WithRunFunc(func() error {
				log.Infof("application running")
				return nil
			}))
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.Contains(t, buf.String(), "[info] application running")
		assert.NotContains(t, string(stdout), "--version version[=true]")
		assert.NotContains(t, string(stdout), "-c, --config FILE")
	})
	t.Run("without run", func(t *testing.T) {
		os.Args = []string{"testApp"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
			jcli.WithLogger(log),
			jcli.DisableConfig(),
			jcli.DisableVersion(),
		)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.NotContains(t, string(stdout), "--version version[=true]")
		assert.NotContains(t, string(stdout), "-c, --config FILE")
		assert.Contains(t, buf.String(), "Starting simple")
	})
	t.Run("help message should contain sub commands", func(t *testing.T) {
		os.Args = []string{"testApp", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
		)
		app.AddCommands(
			jcli.NewCommand("sub1", "sub1 command description",
				jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
					fmt.Println("sub1 command running")
					return nil
				}),
				jcli.WithCommandDesc("sub1 long desc"),
			),
			jcli.NewCommand("sub2", "sub2 command description",
				jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
					fmt.Println("sub2 command running")
					return nil
				}),
			),
			jcli.NewCommand("sub3", "sub3 command description"),
		)
		app.AddCobraCommands(&cobra.Command{
			Use:   "sub4",
			Short: "sub4 command description",
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Println("sub4 command running")
				return nil
			},
		})
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.Contains(t, string(stdout), "sub1 command description")
		assert.Contains(t, string(stdout), "sub2 command description")
		assert.Contains(t, string(stdout), "sub3 command description")
		assert.Contains(t, string(stdout), "sub4 command description")
		assert.Contains(t, string(stdout), "--username")
		assert.Contains(t, string(stdout), "--password")
		assert.Contains(t, string(stdout), "--version version[=true]")
		assert.Contains(t, string(stdout), "-c, --config FILE")
	})

	t.Run("help message should contain sub-sub commands", func(t *testing.T) {
		os.Args = []string{"testApp", "sub1", "--help"}
		r, w, _ := os.Pipe()
		tmp := os.Stdout
		defer func() {
			os.Stdout = tmp
		}()
		os.Stdout = w
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithBaseName("testApp"),
		)
		sub1 := jcli.NewCommand("sub1", "sub1 command description",
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				fmt.Println("sub1 command running")
				return nil
			}),
			jcli.WithCommandDesc("sub1 long desc"),
		)
		sub1.AddCommands(jcli.NewCommand("sub1-sub1", "sub1-sub1 command description"))
		app.AddCommands(sub1)
		app.Run()
		_ = w.Close()
		stdout, _ := io.ReadAll(r)
		assert.Contains(t, string(stdout), "sub1-sub1 command description")
		assert.NotContains(t, string(stdout), "--username")
		assert.NotContains(t, string(stdout), "--password")
		assert.NotContains(t, string(stdout), "--version version[=true]")
		assert.NotContains(t, string(stdout), "-c, --config FILE")
	})

	t.Run("should run sub commands", func(t *testing.T) {
		os.Args = []string{"testApp", "sub1"}
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithLogger(log),
			jcli.WithBaseName("testApp"),
		)
		app.AddCommands(
			jcli.NewCommand("sub1", "sub1 command description",
				jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
					log.Infof("sub1 command running")
					return nil
				}),
				jcli.WithCommandDesc("sub1 long desc"),
			),
		)

		app.Run()
		assert.Contains(t, buf.String(), "sub1 command running")
	})

	t.Run("should run sub-sub commands", func(t *testing.T) {
		os.Args = []string{"testApp", "sub1", "sub1-sub1"}
		var buf bytes.Buffer
		log := newTestLogger(&buf)
		app := jcli.New("simple",
			jcli.WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			jcli.WithLogger(log),
			jcli.WithBaseName("testApp"),
		)
		sub1 := jcli.NewCommand("sub1", "sub1 command description",
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				log.Infof("sub1 command running")
				return nil
			}),
			jcli.WithCommandDesc("sub1 long desc"),
		)
		sub1.AddCommands(
			jcli.NewCommand(
				"sub1-sub1",
				"sub1-sub1 command description",
				jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
					log.Infof("sub1-sub1 command running")
					return nil
				}),
			),
		)
		app.AddCommands(sub1)

		app.Run()
		assert.Contains(t, buf.String(), "sub1-sub1 command running")
	})
}
