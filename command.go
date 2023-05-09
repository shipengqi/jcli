package jcli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RunCommandFunc defines the application's command startup callback function.
type RunCommandFunc func(args []string) error

// Command is a sub command structure of a cli application.
// It is recommended that a command be created with the app.NewCommand()
// function.
type Command struct {
	name    string
	short   string
	desc    string
	aliases []string
	opts    CliOptions
	subs    []*cobra.Command
	runfunc RunCommandFunc
}

// NewCommand creates a new sub command instance based on the given command name
// and other options.
func NewCommand(name string, short string, opts ...CommandOption) *Command {
	c := &Command{
		name:  name,
		short: short,
	}
	c.withOptions(opts...)

	return c
}

// AddCommands adds multiple sub commands to the current command.
func (c *Command) AddCommands(commands ...*Command) {
	for _, v := range commands {
		c.subs = append(c.subs, v.cobraCommand())
	}
}

// AddCobraCommands adds multiple sub cobra.Command to the current command.
func (c *Command) AddCobraCommands(commands ...*cobra.Command) {
	c.subs = append(c.subs, commands...)
}

func (c *Command) cobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     c.name,
		Short:   c.short,
		Aliases: c.aliases,
	}
	if c.desc == "" {
		cmd.Long = c.short
	}

	cmd.Flags().SortFlags = false
	if len(c.subs) > 0 {
		cmd.AddCommand(c.subs...)
	}
	if c.runfunc != nil {
		cmd.Run = c.run
	}
	if c.opts != nil {
		for _, f := range c.opts.Flags().FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
	}

	return cmd
}

func (c *Command) run(cmd *cobra.Command, args []string) {
	if c.runfunc != nil {
		if err := c.runfunc(args); err != nil {
			fmt.Printf("%v %v\n", Red("Error:"), err)
			os.Exit(1)
		}
	}
}

// withOptions apply options for the application.
func (c *Command) withOptions(opts ...CommandOption) *Command {
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}
