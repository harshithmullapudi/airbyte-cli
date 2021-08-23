package airbyte

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/jedib0t/go-pretty/v6/table"
)

//Sources Print
func PrintSourcesTable(sources models.Sources) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Source Definition Id", "Source Id", "Name", "Source Name"})
	for index, s := range sources {
		t.AppendRow([]interface{}{index + 1, s.SourceDefinitionId, s.SourceId, s.Name, s.SourceName})
	}
	t.Render()
}

func PrintSourceTable(source models.Source) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Source Definition Id", "Source Id", "Name", "Source Name"})

	t.AppendRow([]interface{}{1, source.SourceDefinitionId, source.SourceId, source.Name, source.SourceName})

	t.Render()
}

func PrintSources(sources models.Sources) {

	b, err := json.MarshalIndent(sources, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

func PrintSource(source models.Source) {

	b, err := json.MarshalIndent(source, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

// Print Destinations
func PrintDestinationsTable(destinations models.Destinations) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Destination Definition Id", "Destination Id", "Name", "Destination Name"})
	for index, d := range destinations {
		t.AppendRow([]interface{}{index + 1, d.DestinationDefinitionId, d.DestinationId, d.Name, d.DestinationName})
	}
	t.Render()
}

func PrintDestinationTable(destination models.Destination) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Destination Definition Id", "Destination Id", "Name", "Destination Name"})

	t.AppendRow([]interface{}{1, destination.DestinationDefinitionId, destination.DestinationId, destination.Name, destination.DestinationName})

	t.Render()
}

func PrintDestinations(destinations models.Destinations) {

	b, err := json.MarshalIndent(destinations, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

func PrintDestination(destination models.Destination) {

	b, err := json.MarshalIndent(destination, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

//Connections print
func PrintConnectionsTable(connections models.Connections) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Connection Id", "Source Id", "Name", "Source Name", "Destination Id", "Schedule", "Status", "Sync Status"})
	for index, c := range connections {
		t.AppendRow([]interface{}{index + 1, c.ConnectionId, c.SourceId, c.Source.Name, c.Source.SourceName, c.DestinationId, c.Schedule, c.Status, c.LatestSyncJobStatus})
	}
	t.Render()
}

func PrintConnectionTable(connection models.Connection) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Connection Id", "Source Id", "Name", "Source Name", "Destination Id", "Schedule", "Status", "Sync Status"})
	t.AppendRow([]interface{}{1, connection.ConnectionId, connection.SourceId, connection.Source.Name, connection.Source.SourceName, connection.DestinationId, connection.Schedule, connection.Status, connection.LatestSyncJobStatus})
	t.Render()
}

func PrintConnections(connections models.Connections) {

	b, err := json.MarshalIndent(connections, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

func PrintConnection(connection models.Connection) {

	b, err := json.MarshalIndent(connection, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

// Print Jobs
func PrintJobsTable(jobs models.Jobs) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Job Id", "Config Id", "Config Type", "Created At", "Status", "Total Attempts", "Records Synced"})
	for index, j := range jobs {
		var attemtStatus models.Attempt
		for _, a := range j.Attempts {
			if a.Status == "succeeded" {
				attemtStatus = a
			}
		}

		unixTimeUTC := time.Unix(j.Job.CreatedAt, 0)
		unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339)

		t.AppendRow([]interface{}{index + 1, j.Job.Id, j.Job.ConfigId, j.Job.ConfigType, unitTimeInRFC3339, j.Job.Status, len(j.Attempts), attemtStatus.RecordsSynced})
	}
	t.Render()
}

func PrintJobTable(job models.GetJobResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Job Id", "Config Id", "Config Type", "Created At", "Status", "Total Attempts", "Records Synced"})
	var attemtStatus models.Attempt
	for _, a := range job.Attempts {
		if a.Attempt.Status == "succeeded" {
			attemtStatus = a.Attempt
		}
	}

	unixTimeUTC := time.Unix(job.Job.CreatedAt, 0)
	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339)

	t.AppendRow([]interface{}{1, job.Job.Id, job.Job.ConfigId, job.Job.ConfigType, unitTimeInRFC3339, job.Job.Status, len(job.Attempts), attemtStatus.RecordsSynced})
	t.Render()
}

// Print Source Definition
func PrintSourceDefinitionsTable(source_definitions models.SourceDefinitions) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Source Definition Id", "Name", "Docker Repository", "Docker Image Tag"})
	for index, sd := range source_definitions {
		t.AppendRow([]interface{}{index + 1, sd.SourceDefinitionId, sd.Name, sd.DockerRepository, sd.DockerImageTag})
	}
	t.Render()
}

//Print Workspaces
func PrintWorkspacesTable(workspaces models.Workspaces) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Workspace ID", "Workspace Name", "Email", "Slug"})
	for index, w := range workspaces {
		t.AppendRow([]interface{}{index + 1, w.WorkspaceId, w.Name, w.Email, w.Slug})
	}
	t.Render()
}

func PrintWorkspaceTable(workspace models.Workspace) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Workspace ID", "Workspace Name", "Email", "Slug"})

	t.AppendRow([]interface{}{1, workspace.WorkspaceId, workspace.Name, workspace.Email, workspace.Slug})

	t.Render()
}

func PrintWorkspaces(workspaces models.Workspaces) {

	b, err := json.MarshalIndent(workspaces, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

func PrintWorkspace(workspace models.Workspace) {

	b, err := json.MarshalIndent(workspace, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}
