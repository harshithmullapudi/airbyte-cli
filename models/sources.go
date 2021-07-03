package models

type Source struct {
	SourceDefinitionId      string
	SourceId                string
	WorkspaceId             string
	ConnectionConfiguration map[string]interface{}
	Name                    string
	SourceName              string
}

type Sources []Source

type SourceResponse struct {
	Sources []Source
}
