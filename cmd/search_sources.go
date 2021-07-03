package cmd

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/spf13/cobra"
)

var SourcesSearchCmd = &cobra.Command{
	Use:   "sources",
	Args:  cobra.MinimumNArgs(1),
	Short: "Search sources",
	Long:  `Search all sources.`,
	Run: func(cmd *cobra.Command, args []string) {
		var searchString string = args[0]

		sources, err := airbyte.SearchSource(searchString)

		if err != nil {
			cobra.CheckErr(err)
			return
		}

		airbyte.PrintSourcesTable(sources)
	},
}
