package get

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var JobsSubCmd = &cobra.Command{
	Use:   "jobs [configId] [configType]",
	Short: "Get jobs",
	Long: `Fetch all jobs with pagination.

	You can use page(p) and offset(o) to fetch sources respectively`,
	Run: func(cmd *cobra.Command, args []string) {
		var configId string = args[0]

		format, _ := cmd.Flags().GetString("format")
		configType, _ := cmd.Flags().GetString("type")

		var jobs models.Jobs
		jobs, err := airbyte.PaginateJobs(configId, configType)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintJobsTable(jobs)
		} else {
		}
	},
}

func init() {
	JobsSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")
	JobsSubCmd.PersistentFlags().StringP("type", "t", "sync", "Config Type")
}
