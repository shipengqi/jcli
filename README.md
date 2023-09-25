# jcli

A package for building Go applications.

[![Test](https://github.com/shipengqi/jcli/actions/workflows/test.yaml/badge.svg)](https://github.com/shipengqi/jcli/actions/workflows/test.yaml)
[![Codecov](https://codecov.io/gh/shipengqi/jcli/branch/main/graph/badge.svg)](https://codecov.io/gh/shipengqi/jcli)
[![Go Report Card](https://goreportcard.com/badge/github.com/shipengqi/jcli)](https://goreportcard.com/report/github.com/shipengqi/jcli)
[![Release](https://img.shields.io/github/release/shipengqi/jcli.svg)](https://github.com/shipengqi/jcli/releases)
[![License](https://img.shields.io/github/license/shipengqi/jcli)](https://github.com/shipengqi/jcli/blob/main/LICENSE)

## Getting Started

```go
package main

import (
	"fmt"

	"github.com/shipengqi/jcli"
	"github.com/spf13/cobra"
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

func main() {
	// Create a new App instance
	app := jcli.New("demo",
		jcli.WithCliOptions(&fakeCliOptions{}),
		jcli.WithBaseName("demo"),
		jcli.WithExamples("This is a example for demo"),
		jcli.WithDesc("This is a description for demo"),
		jcli.WithAliases("alias1", "alias2"),
		jcli.WithRunFunc(func() error {
			fmt.Println("application running")
			return nil
		}))
	// Add sub commands
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
	// Add cobra commands as sub commands
	app.AddCobraCommands(&cobra.Command{
		Use:   "sub4",
		Short: "sub4 command description",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("sub4 command running")
			return nil
		},
	})
	app.Run()
}
```

### WithLogger

Use `jcli.WithLogger` to set a custom `Logger`

### DisableConfig

By default, `App` will add the `--config` flag, and use [Viper](https://github.com/spf13/viper) to parse the config file.
You can use `DisableConfig` to disable it.

### DisableVersion

By default, `App` will add the `--version` flag, you can use `DisableVersion` to disable it.

### WithSilence 

Use `WithSilence` to set the application to silent mode.

### WithOnSignalReceived 

Use `WithOnSignalReceived` to set a signals' receiver. `SIGTERM` and `SIGINT` are registered by default.
Register other signals via the signal parameter.

### EnableCompletion

Use `EnableCompletion` to create a default 'completion' command.

### Create a new root command

```go
package main

import (
	"fmt"

	"github.com/shipengqi/jcli"
	"github.com/spf13/cobra"
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

func main() {
	// Create a root Command
	app := jcli.NewCommand(
		"demo",
		"This is a short description",
		jcli.WithCommandDesc("This is a long description"),
		jcli.EnableCommandVersion(), // Enable the version flag of the root Command, set only when use the Command as a root command.
		jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
			fmt.Printf("%s Version: \n%s", "=======>", "dev")
			return cmd.Help()
		}),
	)

	// Set PersistentPreRun for the root command
	app.CobraCommand().PersistentPreRun = func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun")
	}
	// Add sub commands
	app.AddCommands(
		jcli.NewCommand("sub1", "sub1 command description",
			jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
				fmt.Println("sub1 command running")
				return nil
			}),
			jcli.WithCommandDesc("sub1 long desc"),
		),
	)
	// Add cobra commands as sub commands
	app.AddCobraCommands(&cobra.Command{
		Use:   "sub2",
		Short: "sub2 command description",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("sub2 command running")
			return nil
		},
	})
	cobra.EnableCommandSorting = false
}
```

## Documentation

You can find the docs at [go docs](https://pkg.go.dev/github.com/shipengqi/jcli).

## 🔋 JetBrains OS licenses

`jcli` had been being developed with **GoLand** under the **free JetBrains Open Source license(s)** granted by JetBrains s.r.o., hence I would like to express my thanks here.

<a href="https://www.jetbrains.com/?from=jcli" target="_blank"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo." width="250" align="middle"></a>
