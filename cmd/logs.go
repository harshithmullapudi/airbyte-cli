package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/harshithmullapudi/airbyte/airbyte/api"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [jobId]",
	Short: "Fetch logs for a job",
	Long: `Fetch logs for a job
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cobra.CheckErr(errors.New("Missing job ID"))
		}

		var jobId int
		jobId, _ = strconv.Atoi(args[0])

		attemptNumber, _ := cmd.Flags().GetInt("attempt")
		job, err := api.GetJob(jobId)
		if err != nil {
			cobra.CheckErr(err)
		}

		for index, attempt := range job.Attempts {
			if (index + 1) == attemptNumber {
				for _, logs := range attempt.Logs.LogLines {
					fmt.Println(logs)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.PersistentFlags().IntP("attempt", "a", 1, "Attempt number")
}
