package jcli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/shipengqi/golib/sysutil"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ConfigFlagName = "config"

var _filename string

func init() {
	pflag.StringVarP(&_filename, ConfigFlagName, "c", _filename,
		"Read configuration from specified `FILE`, support JSON, TOML, YAML, HCL, or Java properties formats.")
}

// addConfigFlag adds flags for a specific server to the specified FlagSet
// object.
func (a *App) addConfigFlag(basename string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(ConfigFlagName))

	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.ReplaceAll(strings.ToUpper(basename), "-", "_"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	cobra.OnInitialize(func() {
		if _filename != "" {
			viper.SetConfigFile(_filename)
		} else {
			viper.AddConfigPath(".")

			if names := strings.Split(basename, "-"); len(names) > 1 {
				viper.AddConfigPath(filepath.Join(sysutil.HomeDir(), "."+names[0]))
				viper.AddConfigPath(filepath.Join("/etc", names[0]))
			}

			viper.SetConfigName(basename)
		}

		if err := viper.ReadInConfig(); err != nil {
			var configFileNotFoundError viper.ConfigFileNotFoundError
			// cannot find any configuration files in the /home and /etc directories
			if _filename == "" && errors.As(err, &configFileNotFoundError) {
				return
			}
			a.logger.Fatalf("Error: failed to read configuration file(%s): %v\n", _filename, err)
		}
	})
}

func (a *App) PrintWorkingDir() {
	wd, _ := os.Getwd()
	a.logger.Infof("%v WorkingDir: %s", progressMessage, wd)
}
