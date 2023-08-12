package jcli

import (
	"fmt"
	"os"

	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/term"
	"github.com/shipengqi/component-base/version/verflag"
	"github.com/shipengqi/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// RunCommandFunc defines the application's command startup callback function.
type RunCommandFunc func(cmd *Command, args []string) error

// Command is a sub command structure of a cli application.
// It is recommended that a command be created with the app.NewCommand()
// function.
type Command struct {
	name          string
	short         string
	desc          string
	examples      string
	enableVersion bool
	aliases       []string
	opts          CliOptions
	subs          []*cobra.Command
	cmd           *cobra.Command
	runfunc       RunCommandFunc
}

// NewCommand creates a new sub command instance based on the given command name
// and other options.
func NewCommand(name string, short string, opts ...CommandOption) *Command {
	c := &Command{
		name:  name,
		short: short,
	}
	c.withOptions(opts...)

	c.cmd = c.cobraCommand()
	return c
}

// AddCommands adds multiple sub commands to the current command.
func (c *Command) AddCommands(commands ...*Command) {
	for _, v := range commands {
		c.subs = append(c.subs, v.cobraCommand())
		c.cmd.AddCommand(v.cobraCommand())
	}
}

// AddCobraCommands adds multiple sub cobra.Command to the current command.
func (c *Command) AddCobraCommands(commands ...*cobra.Command) {
	c.subs = append(c.subs, commands...)
	c.cmd.AddCommand(commands...)
}

// CobraCommand returns cobra command instance inside the Command.
func (c *Command) CobraCommand() *cobra.Command {
	return c.cobraCommand()
}

// Help runs the Help of cobra command.
func (c *Command) Help() error {
	if c.cmd == nil {
		return nil
	}
	return c.cmd.Help()
}

// Name returns the command's name: the first word in the use line.
func (c *Command) Name() string {
	if c.cmd == nil {
		return ""
	}
	return c.cmd.Name()
}

// Flags returns the complete FlagSet that applies
// to this command.
func (c *Command) Flags() *pflag.FlagSet {
	if c.cmd == nil {
		return nil
	}
	return c.cmd.Flags()
}

// MarkHidden sets flags to 'hidden' in your program.
func (c *Command) MarkHidden(flags ...string) {
	if c.cmd == nil {
		return
	}
	for _, v := range flags {
		_ = c.cmd.Flags().MarkHidden(v)
	}
}

// Run runs the command.
func (c *Command) Run() {
	if c.cmd == nil {
		return
	}
	if err := c.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", Red("Error:"), err)
		os.Exit(1)
	}
}

func (c *Command) cobraCommand() *cobra.Command {
	if c.cmd != nil {
		return c.cmd
	}

	cmd := &cobra.Command{
		Use:     c.name,
		Short:   c.short,
		Aliases: c.aliases,
		Example: c.examples,
		// stop printing usage when the command errors
		// ensure that silence can take effect when the command is separated from the application.
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	if c.desc == "" {
		cmd.Long = c.short
	} else {
		cmd.Long = c.desc
	}

	cmd.Flags().SortFlags = false
	if len(c.subs) > 0 {
		cmd.AddCommand(c.subs...)
	}

	cmd.RunE = c.run

	var nfs cliflag.NamedFlagSets
	if c.opts != nil {
		nfs = c.opts.Flags()
		for _, f := range nfs.FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
	}
	addHelpCommandFlag(c.name, cmd.Flags())

	// add version flag when the command is separated from the application.
	if c.enableVersion {
		verflag.AddFlags(nfs.FlagSet(FlagSetNameGlobal))
	}

	width, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, nfs, width)
	return cmd
}

func (c *Command) run(_ *cobra.Command, args []string) error {
	if c.enableVersion {
		verflag.PrintAndExitIfRequested()
	}

	if c.opts != nil {
		if err := c.applyOptions(); err != nil {
			return err
		}
	}
	if c.runfunc != nil {
		if err := c.runfunc(c, args); err != nil {
			return err
		}
	}
	return nil
}

// withOptions apply options for the application.
func (c *Command) withOptions(opts ...CommandOption) *Command {
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

func (c *Command) applyOptions() error {
	if options, ok := c.opts.(CompletableOptions); ok {
		if err := options.Complete(); err != nil {
			return err
		}
	}

	if errs := c.opts.Validate(); len(errs) != 0 {
		return errors.NewAggregate(errs)
	}

	return nil
}
