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
package create

import (
	"errors"
	"os"

	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/spf13/cobra"
)

// checkCmd represents the get command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Sources/Destinations/Connections",
	Long: `Create Sources/Destinations/Connections from a config folder. You need to have SOURCE_CONNECTION.yaml(sources), DESTINATION_CONNECTION.yaml(destinations), STANDARD_SYNC.yaml(connections) files.
	Note: This will neglect if the entity is already created.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		configFolder, _ := cmd.Flags().GetString("folder")
		create, _ := cmd.Flags().GetBool("create")

		if configFolder == "" {
			logger.Error("provide config folder")
			cobra.CheckErr(errors.New("provide config folder"))
		}

		// Check if config folder exist
		_, err := os.Stat(configFolder)
		if os.IsNotExist(err) {
			logger.Error("No config folder found")
			cobra.CheckErr(err)
		}

		// Start with sources

		logger.Debug("Starting sources creation")
		CreateSources(configFolder, create)

		// Create destinations
		logger.Debug("Starting destinations creation")
		CreateDestinations(configFolder, create)

		// Create Connections
		logger.Debug("Starting connections creation")
		CreateConnections(configFolder, create)
	},
}

func init() {
	CreateCmd.PersistentFlags().StringP("folder", "f", "", "Config folder")
	CreateCmd.PersistentFlags().BoolP("create", "c", false, "Setting this to false will only validate sources and doesn't create")
}
