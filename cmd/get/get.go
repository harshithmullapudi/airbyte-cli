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
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [sub]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Get details about sources, destinations, connections",
	Long:  `Get details about sources, destinations, connections`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Notice("Kindly specify resource. Example (sources, connections, destinations, source, destination, connection)")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// sub commands for get
	getCmd.AddCommand(SourcesSubCmd)
	getCmd.AddCommand(DestinationsSubCmd)
	getCmd.AddCommand(ConnectionsSubCmd)
	getCmd.AddCommand(SourceSubCmd)
	getCmd.AddCommand(ConnectionSubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
