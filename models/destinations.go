package models

type DestinationDefinition struct {
	DestinationDefinitionId       string
	SupportedDestinationSyncModes []string
	SupportsDbt                   bool
	SupportsNormalization         bool
}

type Destination struct {
	DestinationDefinitionId string                 `yaml:"destinationDefinitionId"`
	DestinationId           string                 `yaml:"destinationId"`
	WorkspaceId             string                 `yaml:"workspaceId"`
	ConnectionConfiguration map[string]interface{} `yaml:"configuration"`
	Name                    string                 `yaml:"name"`
	DestinationName         string                 `yaml:"destinationName"`
}

type Destinations []Destination

type DestinationResponse struct {
	Destinations []Destination
}
