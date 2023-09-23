package jcli

import (
	"fmt"
	"os"
	"strings"

	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/cli/globalflag"
	"github.com/shipengqi/component-base/term"
	"github.com/shipengqi/component-base/version"
	"github.com/shipengqi/component-base/version/verflag"
	"github.com/shipengqi/errors"
	"github.com/shipengqi/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagSetNameGlobal = "global"
)

var (
	progressMessage = Green("==>")
)

// RunFunc defines the application's run callback function.
type RunFunc func() error

// App is the main structure of a cli application.
type App struct {
	name           string
	basename       string
	description    string
	examples       string
	aliases        []string
	runfunc        RunFunc
	signalReceiver SignalReceiver
	signals        []os.Signal
	setonce        chan struct{}
	sigc           chan os.Signal
	opts           CliOptions
	logger         Logger
	versionLogger  *infoLogger
	silence        bool
	disableVersion bool
	disableConfig  bool
	subs           []*cobra.Command
	cmd            *cobra.Command
}

// New create a new cli application.
func New(name string, opts ...Option) *App {
	a := &App{
		name:    name,
		setonce: make(chan struct{}),
	}
	a.withOptions(opts...)

	// set default logger
	if a.logger == nil {
		a.logger = log.WithValues()
	}
	a.versionLogger = newInfoLogger(a.logger)

	a.cmd = a.buildCommand()

	return a
}

// Run is used to launch the application.
func (a *App) Run() {
	if a.signalReceiver != nil {
		a.setupSignalHandler(a.signalReceiver, a.signals...)
	}

	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", Red("Error:"), err)
		os.Exit(1)
	}
}

// Command returns cobra command instance inside the App.
func (a *App) Command() *cobra.Command {
	return a.cmd
}

// AddCommands adds multiple sub commands to the App.
func (a *App) AddCommands(commands ...*Command) {
	for _, v := range commands {
		// Todo force to remove global version flag for the sub commands
		a.subs = append(a.subs, v.cobraCommand())
		a.cmd.AddCommand(v.cobraCommand())
	}
}

// AddCobraCommands adds multiple sub cobra.Command to the App.
func (a *App) AddCobraCommands(commands ...*cobra.Command) {
	a.subs = append(a.subs, commands...)
	a.cmd.AddCommand(commands...)
}

// Logger return the (logger) of the App.
func (a *App) Logger() Logger {
	return a.logger
}

// withOptions apply options for the application.
func (a *App) withOptions(opts ...Option) *App {
	for _, opt := range opts {
		opt.apply(a)
	}
	return a
}

func (a *App) buildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     NormalizeCliName(a.basename),
		Short:   a.name,
		Aliases: a.aliases,
		Long:    a.description,
		Example: a.examples,
		// stop printing usage when the command errors
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	cliflag.InitFlags(cmd.Flags())

	if len(a.subs) > 0 {
		cmd.AddCommand(a.subs...)
	}
	cmd.SetHelpCommand(helpCommand(NormalizeCliName(a.basename)))
	// Todo make it configurable EnableCompletion(hidden bool), add it for Command as well?? , add tests
	cmd.CompletionOptions.DisableDefaultCmd = true
	// always add App.run func
	cmd.RunE = a.run

	var nfs cliflag.NamedFlagSets

	// Add command line flag sets
	if a.opts != nil {
		nfs = a.opts.Flags()
		fs := cmd.Flags()
		for _, set := range nfs.FlagSets {
			fs.AddFlagSet(set)
		}
	}

	cmd.Flags().AddFlagSet(nfs.FlagSet(FlagSetNameGlobal))
	if !a.disableVersion {
		// add version flag
		verflag.AddFlags(nfs.FlagSet(FlagSetNameGlobal))
	}
	if !a.disableConfig {
		a.addConfigFlag(a.basename, nfs.FlagSet(FlagSetNameGlobal))
	}
	globalflag.AddGlobalFlags(nfs.FlagSet(FlagSetNameGlobal), cmd.Name())

	width, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, nfs, width)

	return cmd
}

func (a *App) run(cmd *cobra.Command, args []string) error {
	if !a.disableVersion {
		verflag.PrintAndExitIfRequested()
	}

	if !a.silence {
		a.PrintWorkingDir()
	}
	cliflag.PrintFlags(cmd.Flags(), a.versionLogger)

	if !a.disableConfig && a.opts != nil {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		if err := viper.Unmarshal(a.opts); err != nil {
			return err
		}
	}

	if !a.silence {
		a.logger.Infof("%s Starting %s ...", progressMessage, a.name)
		a.logger.Infof("%s Executed %s ...", progressMessage, strings.Join(args, " "))
		if !a.disableVersion {
			a.logger.Infof("%s Version: \n%s", progressMessage, version.Get().String())
		}
		if !a.disableConfig && viper.ConfigFileUsed() != "" {
			a.logger.Infof("%s Config file used: `%s`", progressMessage, viper.ConfigFileUsed())
		}
	}

	if a.opts != nil {
		if err := a.applyOptions(); err != nil {
			return err
		}
	}

	if a.runfunc != nil {
		return a.runfunc()
	}

	return nil
}

func (a *App) applyOptions() error {
	if options, ok := a.opts.(CompletableOptions); ok {
		if err := options.Complete(); err != nil {
			return err
		}
	}

	if errs := a.opts.Validate(); len(errs) != 0 {
		return errors.NewAggregate(errs)
	}

	if options, ok := a.opts.(PrintableOptions); ok && !a.silence {
		a.logger.Infof("%s Options: `%s`", progressMessage, options.String())
	}

	return nil
}
