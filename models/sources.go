package models

type SourceDefinition struct {
	SourceDefinitionId string
	Name               string
	DockerRepository   string
	DockerImageTag     string
	DocumentationUrl   string
	Icon               string
}

type SourceDefinitions []SourceDefinition
type Source struct {
	SourceDefinitionId      string
	SourceId                string
	WorkspaceId             string
	ConnectionConfiguration map[string]interface{}
	Name                    string
	SourceName              string
	SourceDefinition        SourceDefinition
}

type Sources []Source

type SourceResponse struct {
	Sources Sources
}

type SourceDefinitionResponse struct {
	SourceDefinitions SourceDefinitions
}

type SourceCheckResponse struct {
	Status  string
	Message string
	JobInfo JobInfo
}
