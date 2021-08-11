package get

import (
	"strconv"

	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var JobsSubCmd = &cobra.Command{
	Use:   "jobs [configId] [configType]",
	Short: "Returns recent jobs for a connection. Jobs are returned in descending order by createdAt",
	Long:  `Fetch all jobs. Check https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/jobs/list`,
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

var JobSubCmd = &cobra.Command{
	Use:   "job [jobId]",
	Short: "Get information about a job",
	Long:  `Get information about a job. Check https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/jobs/get`,
	Run: func(cmd *cobra.Command, args []string) {
		var jobId int
		jobId, _ = strconv.Atoi(args[0])

		format, _ := cmd.Flags().GetString("format")

		job, err := airbyte.FetchJob(jobId)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintJobTable(job)
		} else {
		}
	},
}

func init() {
	JobsSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")
	JobsSubCmd.PersistentFlags().StringP("type", "t", "sync", "Config Type")

	JobSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")
}
