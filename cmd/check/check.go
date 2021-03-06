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
package check

import (
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/spf13/cobra"
)

// checkCmd represents the get command
var CheckCmd = &cobra.Command{
	Use:   "check [sub]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Check connection to Source/Destination",
	Long:  `Validate if the config and secrets are right for Source/Destination`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Notice("Kindly specify resource. Example (source, destination)")
	},
}

func init() {
	// sub commands for get
	CheckCmd.AddCommand(SourceSubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
