package search

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/spf13/cobra"
)

var ConnectionsSearchCmd = &cobra.Command{
	Use:   "connections",
	Args:  cobra.MinimumNArgs(1),
	Short: "Search connections",
	Long:  `Search all connections.`,
	Run: func(cmd *cobra.Command, args []string) {
		var searchString string = args[0]

		connections, err := airbyte.SearchConnection(searchString)

		if err != nil {
			cobra.CheckErr(err)
			return
		}

		airbyte.PrintConnectionsTable(connections)
	},
}
