/*
Copyright © 2019 Rael Garcia <rael@rael.io>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var verbose bool
var taskFile string
var cfgFile string

const defaultConfigDir = ".yodo"
const defaultConfigFile = ".yodo"
const defaultConfigFormat = "yaml"

const defaultList = "default"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yodo",
	Short: "Simple TODO list application, to learn Cobra and Viper",
	Long: `Simple TODO list command-line application written in GoLang using Cobra and Viper libraries.
This application is just a learning exercise.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yodo.yaml)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", viper.GetBool("verbose"), "verbose")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {

		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

	} else {

		cfgFile = defaultConfigFile

		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType(defaultConfigFormat)
		viper.SetConfigName(defaultConfigFile)

		viper.SetDefault("verbose", false)
		viper.SetDefault("list", defaultList)
		viper.SetDefault("list_dir", filepath.Join(home, defaultConfigDir))
		viper.SetDefault("list_file", fmt.Sprintf("%s/%s.%s", viper.GetString("list_dir"), defaultList, defaultListFormat))
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
