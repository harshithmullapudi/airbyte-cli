package models

type Stream struct {
	Stream map[string]interface{}
	Config map[string]interface{}
}

type SyncCatalog struct {
	Streams []Stream
}

type Schedule struct {
	Units    int64
	TimeUnit string
}

type Connection struct {
	ConnectionId    string
	Name            string
	NamespaceFormat string
	Prefix          string
	SourceId        string
	DestinationId   string
	OperationIds    []string
	SyncCatalog     SyncCatalog
	Schedule        Schedule
	Status          string
	Source          Source
}

type Connections []Connection

type ConnectionResponse struct {
	Connections []Connection
}
