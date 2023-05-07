package jcli

import (
	"fmt"
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"testing"
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

	t.Run("with sort flags", func(t *testing.T) {
		app := New("simple",
			WithCliOptions(&fakeCliOptions{"Pooky", "PASS"}),
			WithBaseName("testApp"),
			WithDesc("test application description"),
			WithSortFlags(false))
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
}
