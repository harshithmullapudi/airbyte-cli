package get

import (
	"github.com/harshithmullapudi/airbyte/airbyte"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/cobra"
)

var WorkspacesSubCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "Return all workspaces",
	Long:  `Return all workspaces with pagination. You can use page(p) and offset(o) to fetch workspaces respectively. Check this https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/workspaces/list`,
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetInt("number")
		offset, _ := cmd.Flags().GetInt("offset")
		format, _ := cmd.Flags().GetString("format")
		var workspaces models.Workspaces
		workspaces, err := airbyte.PaginateWorkspaces(offset, number)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintWorkspacesTable(workspaces)
		} else {
			airbyte.PrintWorkspaces(workspaces)
		}
	},
}

var WorkspaceSubCmd = &cobra.Command{
	Use:   "workspace [workspace Id]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Get a workspace",
	Long:  `Get a workspace in table/json format. Check this https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html#post-/v1/workspaces/get`,
	Run: func(cmd *cobra.Command, args []string) {
		var workspaceId string = args[0]
		format, _ := cmd.Flags().GetString("format")
		workspace, err := airbyte.FetchWorkspace(workspaceId)

		if err != nil {
			cobra.CheckErr(err)
		}

		if format == "table" {
			airbyte.PrintWorkspaceTable(workspace)
		} else {
			airbyte.PrintWorkspace(workspace)
		}
	},
}

func init() {
	WorkspacesSubCmd.PersistentFlags().IntP("number", "n", 10, "Number of sources to fetch")
	WorkspacesSubCmd.PersistentFlags().IntP("offset", "o", 0, "Offset")
	WorkspacesSubCmd.PersistentFlags().StringP("format", "f", "table", "Print format")

	WorkspaceSubCmd.PersistentFlags().StringP("format", "f", "json", "Print format")
}
