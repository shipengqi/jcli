package jcli

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"

	cliflag "github.com/shipengqi/component-base/cli/flag"
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
	t.Run("simple", func(t *testing.T) {
		app := New("simple")
		app.Run()
	})

	t.Run("with cli options", func(t *testing.T) {
		app := New("simple", WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}))
		app.Run()
	})

	t.Run("with basename", func(t *testing.T) {
		app := New("simple",
			WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			WithBaseName("testApp"))
		app.Run()
	})

	t.Run("with run", func(t *testing.T) {
		app := New("simple",
			WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			WithBaseName("testApp"),
			WithDesc("test application description"),
			WithRunFunc(func() error {
				fmt.Println("application running")
				return nil
			}),
			DisableConfig())
		app.Run()
	})

	t.Run("with silence", func(t *testing.T) {
		app := New("simple",
			WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			WithBaseName("testApp"),
			WithDesc("test application description"),
			WithRunFunc(func() error {
				fmt.Println("application running")
				return nil
			}),
			DisableConfig(),
			WithSilence())
		app.Run()
	})

	t.Run("disable version and config flags", func(t *testing.T) {
		app := New("simple",
			WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			WithBaseName("testApp"),
			WithDesc("test application description"),
			DisableConfig(),
			DisableVersion())
		app.Run()
	})

	t.Run("with run with config", func(t *testing.T) {
		app := New("simple",
			WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			WithBaseName("testApp"),
			WithDesc("test application description"),
			WithRunFunc(func() error {
				fmt.Println("application running")
				return nil
			}))
		app.Run()
	})

	t.Run("add commands", func(t *testing.T) {
		app := New("simple",
			WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			WithBaseName("testApp"),
			WithDesc("test application description"),
			DisableCommandSorting(),
			DisableConfig(),
		)

		app.AddCommands(
			NewCommand("sub1", "sub1 command description", WithCommandRunFunc(func(args []string) error {
				fmt.Println("sub1 command running")
				return nil
			})),
			NewCommand("sub2", "sub2 command description", WithCommandRunFunc(func(args []string) error {
				fmt.Println("sub2 command running")
				return nil
			})),
		)
		app.AddCobraCommands(&cobra.Command{
			Use:   "sub3",
			Short: "sub3 command description",
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Println("sub3 command running")
				return nil
			},
		})
		app.Run()
	})
}
