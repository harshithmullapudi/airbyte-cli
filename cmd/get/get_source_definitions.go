package get

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/airbyte/api"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var SourceDefinitionsSubCmd = &cobra.Command{
	Use:   "source_definitions",
	Short: "Get source definitions",
	Long:  `Fetch all source definitions.`,
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		var source_definitions models.SourceDefinitions
		source_definitions, err := api.GetSourceDefinitions()

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintSourceDefinitionsTable(source_definitions)
		} else {
		}
	},
}

func init() {

	SourceDefinitionsSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")
}
