package jcli

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"
)

// CliOptions abstracts configuration options for reading parameters from the
// command line.
type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}

// CompletableOptions abstracts options which can be completed.
type CompletableOptions interface {
	Complete() error
}

// PrintableOptions abstracts options which can be printed.
type PrintableOptions interface {
	String() string
}

// ====================================
// Application Options

// Option defines optional parameters for initializing the application
// structure.
type Option interface {
	apply(a *App)
}

// optionFunc wraps a func, so it satisfies the Option interface.
type optionFunc func(*App)

func (f optionFunc) apply(a *App) {
	f(a)
}

// cmdOptionFunc wraps a func, so it satisfies the Option interface.
type cmdOptionFunc func(*Command)

func (f cmdOptionFunc) apply(c *Command) {
	f(c)
}

// WithRunFunc is used to set the application run callback function option.
func WithRunFunc(run RunFunc) Option {
	return optionFunc(func(a *App) {
		a.runfunc = run
	})
}

// WithBaseName is used to set the basename of the cli.
func WithBaseName(basename string) Option {
	return optionFunc(func(a *App) {
		a.basename = basename
	})
}

// WithDesc is used to set the description of the application.
func WithDesc(desc string) Option {
	return optionFunc(func(a *App) {
		a.description = desc
	})
}

// WithCliOptions to open the application's function to read from the command line
// or read parameters from the configuration file.
func WithCliOptions(opts CliOptions) Option {
	return optionFunc(func(a *App) {
		a.opts = opts
	})
}

// WithSilence sets the application to silent mode, in which the program startup
// information, configuration information, and version information are not
// printed in the console.
func WithSilence() Option {
	return optionFunc(func(a *App) {
		a.silence = true
	})
}

// WithAliases sets the application aliases.
func WithAliases(aliases ...string) Option {
	return optionFunc(func(a *App) {
		a.aliases = aliases
	})
}

// DisableVersion disable the version flag.
func DisableVersion() Option {
	return optionFunc(func(a *App) {
		a.disableVersion = true
	})
}

// DisableConfig disable the config flag.
func DisableConfig() Option {
	return optionFunc(func(a *App) {
		a.disableConfig = true
	})
}

// ====================================
// Command Options

// CommandOption defines optional parameters for initializing the command
// structure.
type CommandOption interface {
	apply(a *Command)
}

// WithCommandCliOptions to open the application's function to read from the
// command line.
func WithCommandCliOptions(opts CliOptions) CommandOption {
	return cmdOptionFunc(func(c *Command) {
		c.opts = opts
	})
}

// WithCommandRunFunc is used to set the application's command startup callback
// function option.
func WithCommandRunFunc(run RunCommandFunc) CommandOption {
	return cmdOptionFunc(func(c *Command) {
		c.runfunc = run
	})
}

// WithCommandAliases sets the command aliases.
func WithCommandAliases(aliases ...string) CommandOption {
	return cmdOptionFunc(func(c *Command) {
		c.aliases = aliases
	})
}

// WithCommandDesc sets the command aliases.
func WithCommandDesc(desc string) CommandOption {
	return cmdOptionFunc(func(c *Command) {
		c.desc = desc
	})
}
