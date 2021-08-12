package get

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/airbyte/api"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var ConnectionsSubCmd = &cobra.Command{
	Use:   "connections",
	Short: "Returns all connections with pagination",
	Long:  `Returns all connections with pagination. You can use page(p) and offset(o) to fetch connections respectively. Check this https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/connections/list`,
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetInt("number")
		offset, _ := cmd.Flags().GetInt("offset")
		format, _ := cmd.Flags().GetString("format")
		status, _ := cmd.Flags().GetString("status")

		var connections models.Connections
		connections, err := airbyte.PaginateConnections(offset, number, status)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintConnectionsTable(connections)
		} else {
			airbyte.PrintConnections(connections)
		}
	},
}

var ConnectionSubCmd = &cobra.Command{
	Use:   "connection [connection Id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Get a connection",
	Long:  `Get a connection in table/json format. Check this https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/connections/get`,
	Run: func(cmd *cobra.Command, args []string) {
		var connectionId string = args[0]
		format, _ := cmd.Flags().GetString("format")

		connection, err := api.GetConnection(connectionId)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintConnectionTable(connection)
		} else {
			airbyte.PrintConnection(connection)
		}

	},
}

func init() {
	ConnectionsSubCmd.PersistentFlags().IntP("number", "n", 10, "Number of sources to fetch")
	ConnectionsSubCmd.PersistentFlags().IntP("offset", "o", 0, "Offset")
	ConnectionsSubCmd.PersistentFlags().StringP("status", "s", "", "Print format")
	ConnectionsSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")

	ConnectionSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")
}
