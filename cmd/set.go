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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set-config",
	Short: "Set your airbyte url and workspaceID",
	Long:  `You need to set airbyte url and workspaceID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter airbyte URL here: ")
		URL := bufio.NewScanner(os.Stdin)
		URL.Scan()
		_, error := common.CheckForURL(URL.Text())

		if error != nil {
			cobra.CheckErr(error)
			return
		}

		// Remove trailing slash if any
		api_url := strings.TrimSuffix(URL.Text(), "/")

		viper.Set("airbyte_url", api_url)

		fmt.Print("Enter current workspace: ")

		workspaceId := bufio.NewScanner(os.Stdin)
		workspaceId.Scan()

		var workspace models.Workspace
		workspace, error = airbyte.CheckIfWorkspaceExist(workspaceId.Text())

		if error != nil {
			cobra.CheckErr(error)
			return
		}

		viper.Set("workspace_id", workspaceId.Text())
		// Finally write config to config file
		viper.WriteConfig()

		logger.Notice("We found workspace with name: " + workspace.Name + " and email: " + workspace.Email)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
