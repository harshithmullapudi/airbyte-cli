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
	SourceDefinitionId      string                 `yaml:"sourceDefinitionId"`
	SourceId                string                 `yaml:"sourceId"`
	WorkspaceId             string                 `yaml:"workspaceId"`
	ConnectionConfiguration map[string]interface{} `yaml:"configuration"`
	Name                    string                 `yaml:"name"`
	SourceName              string                 `yaml:"sourceName"`
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

type SourceCheckConfig struct {
	SourceDefinitionId      string
	ConnectionConfiguration map[string]interface{}
}
