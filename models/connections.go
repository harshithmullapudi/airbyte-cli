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
	ConnectionId           string
	Name                   string
	NamespaceFormat        string
	Prefix                 string
	SourceId               string
	DestinationId          string
	OperationIds           []string
	SyncCatalog            SyncCatalog
	Schedule               Schedule
	Status                 string
	Source                 Source
	Destination            Destination
	LatestSyncJobCreatedAt int64
	LatestSyncJobStatus    string
	IsSyncing              bool
}

type ConnectionShort struct {
	ConnectionId    string      `yaml:"connectionId"`
	Name            string      `yaml:"name"`
	NamespaceFormat string      `yaml:"namespaceForm"`
	Prefix          string      `yaml:"prefix"`
	SourceId        string      `yaml:"sourceId"`
	DestinationId   string      `yaml:"destinationId"`
	OperationIds    []string    `yaml:"operationIds"`
	Catalog         interface{} `yaml:"catalog"`
	Manual          bool        `yaml:"manual"`
	Schedule        Schedule    `yaml:"schedule"`
	Status          string      `yaml:"status"`
}

type Connections []Connection

type ConnectionResponse struct {
	Connections []Connection
}

type ConnectionsShort []ConnectionShort
