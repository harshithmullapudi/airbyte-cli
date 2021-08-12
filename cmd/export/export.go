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
package export

import (
	"errors"
	"os"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/spf13/cobra"
)

// checkCmd represents the get command
var ExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export Sources/Destinations/Connections",
	Long:  `Export Sources/Destinations/Connections`,
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")

		if target == "" {
			logger.Error("No target specified")
			cobra.CheckErr(errors.New("no target specified"))
		}

		logger.Info("Downloading the config...")

		err := common.DownloadFile()

		if err != nil {
			logger.Error(err)
		}

		err = common.Untar()

		if err != nil {
			logger.Error(err)
		}

		logger.Info("Downloaded and extracted the config.")

		if err := os.MkdirAll(target, 0755); err != nil {
			cobra.CheckErr(err)
		}

		err = common.CopyConfigToTarget(target)

		if err != nil {
			logger.Error(err)
			cobra.CheckErr(err)
		}

		logger.Notice("Exported successfully")

		logger.Info("Cleaning up...")
		err = common.CleanUp()
		if err != nil {
			logger.Error(err)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	ExportCmd.PersistentFlags().StringP("target", "t", "", "Target folder to export")
}
