package check

import (
	"fmt"

	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var SourceSubCmd = &cobra.Command{
	Use:   "source [source Id]",
	Short: "Check source",
	Long:  `Check whether the source is valid or not`,
	Run: func(cmd *cobra.Command, args []string) {
		var sourceId string = args[0]
		var sourceCheck models.SourceCheckResponse
		sourceCheck, err := airbyte.CheckSourceConnection(sourceId)

		if err != nil {
			cobra.CheckErr(err)
		}

		if sourceCheck.Status == "succeeded" {
			fmt.Println("\033[32m", "Source "+sourceId+" is valid")
		} else {
			logger.Error("This source " + sourceId + " is invalid - " + sourceCheck.Message)
		}
	},
}
