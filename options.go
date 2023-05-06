package jcli

import (
	"github.com/spf13/pflag"
)

// CliOptions abstracts configuration options for reading parameters from the
// command line.
type CliOptions interface {
	// AddFlags adds flags to the given pflag.FlagSet object.
	AddFlags(fs *pflag.FlagSet)
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

type Option interface {
	apply(a *App)
}

// ====================================
// Application Options

// optionFunc wraps a func, so it satisfies the Option interface.
type optionFunc func(*App)

func (f optionFunc) apply(a *App) {
	f(a)
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

// DisableVersion disable the version flag.
func DisableVersion() Option {
	return optionFunc(func(a *App) {
		a.noVersion = true
	})
}

// DisableConfig disable the config flag.
func DisableConfig() Option {
	return optionFunc(func(a *App) {
		a.noConfig = true
	})
}

// BindViper disable the config flag.
func BindViper() Option {
	return optionFunc(func(a *App) {
		a.bindViper = true
	})
}
