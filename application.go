package jcli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/shipengqi/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RunFunc defines the application's run callback function.
type RunFunc func() error

// App is the main structure of a cli application.
type App struct {
	name        string
	basename    string
	description string
	runfunc     RunFunc
	opts        CliOptions
	silence     bool
	noVersion   bool
	noConfig    bool
	bindViper   bool
	subs        []*cobra.Command
	cmd         *cobra.Command
}

// New create a new cli application.
func New(name string, opts ...Option) *App {
	a := &App{
		name: name,
	}
	a.applyOptions(opts...)

	a.cmd = a.buildCommand()

	return a
}

// Run is used to launch the application.
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

// applyOptions apply options for the application.
func (a *App) applyOptions(opts ...Option) *App {
	for _, opt := range opts {
		opt.apply(a)
	}
	return a
}

// AddCommands adds multiple sub commands to the application.
func (a *App) AddCommands(commands ...*cobra.Command) {
	a.subs = append(a.subs, commands...)
}

func (a *App) buildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   NormalizeCliName(a.basename),
		Short: a.name,
		Long:  a.description,
		// stop printing usage when the command errors
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	if len(a.subs) > 0 {
		for i := range a.subs {
			cmd.AddCommand(a.subs[i])
		}
	}

	if a.runfunc != nil {
		cmd.RunE = a.run
	}

	if a.opts != nil {
		a.opts.AddFlags(cmd.Flags())
	}

	addConfigFlag(a.basename, cmd.Flags())

	return cmd
}

func (a *App) run(cmd *cobra.Command) error {

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	if err := viper.Unmarshal(a.opts); err != nil {
		return err
	}

	if a.opts != nil {
		if errs := a.opts.Validate(); len(errs) > 0 {
			return errors.NewAggregate(errs)
		}
	}

	if a.runfunc != nil {
		return a.runfunc()
	}

	return nil
}
