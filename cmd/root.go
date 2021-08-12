/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/harshithmullapudi/airbyte/cmd/check"
	"github.com/harshithmullapudi/airbyte/cmd/create"
	"github.com/harshithmullapudi/airbyte/cmd/export"
	"github.com/harshithmullapudi/airbyte/cmd/get"
	"github.com/harshithmullapudi/airbyte/cmd/search"
	"github.com/harshithmullapudi/airbyte/logger"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "airbyte",
	Short: "CLI tool for Airbyte",
	Long:  `CLI tool for Airbyte`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default searhes for file $HOME/.airbyte.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add sub commands
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(check.CheckCmd)
	rootCmd.AddCommand(export.ExportCmd)
	rootCmd.AddCommand(create.CreateCmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	dirname, err := os.UserHomeDir()

	if err != nil {
		cobra.CheckErr(err)
	}

	_, err = os.Stat(dirname + "/.airbyte.yaml")

	if os.IsNotExist(err) {
		f, e := os.Create(dirname + "/.airbyte.yaml")
		if e != nil {
			cobra.CheckErr(err)
		}
		defer f.Close()
	}

	if err := os.Mkdir(dirname+"/.airbyte", 0755); !os.IsExist(err) {
		logger.Info("Create .airbyte folder")
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".airbyte" (without extension).
		viper.AddConfigPath(dirname)
		viper.SetConfigName(".airbyte")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
