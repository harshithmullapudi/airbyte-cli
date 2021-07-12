package get

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var DestinationsSubCmd = &cobra.Command{
	Use:   "destinations",
	Short: "Fetch all destinations",
	Long: `Fetch all destinations with pagination.

	You can use page(p) and offset(o) to fetch destinations respectively`,
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetInt("number")
		offset, _ := cmd.Flags().GetInt("offset")
		format, _ := cmd.Flags().GetString("format")
		var destinations models.Destinations
		destinations, err := airbyte.PaginateDestinations(offset, number)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintDestinationsTable(destinations)
		} else {
			airbyte.PrintDestinations(destinations)
		}
	},
}

var DestinationSubCmd = &cobra.Command{
	Use:   "destination [destination Id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Get destination details using destination Id",
	Long:  `This will return destination details either in table/json format. Check this https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/destinations/get`,
	Run: func(cmd *cobra.Command, args []string) {
		var destinationId string = args[0]
		format, _ := cmd.Flags().GetString("format")
		destination, err := airbyte.FetchDestination(destinationId)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintDestinationTable(destination)
		} else {
			airbyte.PrintDestination(destination)
		}
	},
}

func init() {
	DestinationsSubCmd.PersistentFlags().IntP("number", "n", 10, "Number of sources to fetch")
	DestinationsSubCmd.PersistentFlags().IntP("offset", "o", 0, "Offset")
	DestinationsSubCmd.PersistentFlags().StringP("format", "f", "table", "Offset")

	DestinationSubCmd.PersistentFlags().StringP("format", "f", "table", "Offset")
}
