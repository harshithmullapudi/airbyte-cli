package get

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var SourcesSubCmd = &cobra.Command{
	Use:   "sources",
	Short: "Get sources",
	Long: `Fetch all sources with pagination.

You can use page(p) and offset(o) to fetch sources respectively`,
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetInt("number")
		offset, _ := cmd.Flags().GetInt("offset")
		format, _ := cmd.Flags().GetString("format")
		var sources models.Sources
		sources, err := airbyte.PaginateSources(offset, number)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintSourcesTable(sources)
		} else {
			airbyte.PrintSources(sources)
		}
	},
}

var SourceSubCmd = &cobra.Command{
	Use:   "source [source Id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Get source details using source Id",
	Long:  `This will return source details either in table/json format. Check this https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/sources/get`,
	Run: func(cmd *cobra.Command, args []string) {
		var sourceId string = args[0]
		format, _ := cmd.Flags().GetString("format")
		source, err := airbyte.FetchSource(sourceId)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintSourceTable(source)
		} else {
			airbyte.PrintSource(source)
		}
	},
}

func init() {
	SourcesSubCmd.PersistentFlags().IntP("number", "n", 10, "Number of sources to fetch")
	SourcesSubCmd.PersistentFlags().IntP("offset", "o", 0, "Offset")
	SourcesSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")

	SourceSubCmd.PersistentFlags().StringP("format", "f", "json", "Print format")
}
