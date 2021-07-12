package models

type DestinationDefinition struct {
	DestinationDefinitionId       string
	SupportedDestinationSyncModes []string
	SupportsDbt                   bool
	SupportsNormalization         bool
}

type Destination struct {
	DestinationDefinitionId string
	DestinationId           string
	WorkspaceId             string
	ConnectionConfiguration map[string]interface{}
	Name                    string
	DestinationName         string
}

type Destinations []Destination

type DestinationResponse struct {
	Destinations []Destination
}
