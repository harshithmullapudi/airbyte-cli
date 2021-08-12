package api

import (
	"encoding/json"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/viper"
)

// Get Connections - /api/v1/web_backend/connections/list
func GetConnections() (models.Connections, error) {
	var API_URL string = common.GetFullApiURL(GET_CONNECTIONS)

	workspaceId := viper.GetString("workspace_id")

	respBody, err := common.ApiCall(API_URL, map[string]string{
		"workspaceId": workspaceId,
	})

	//Handle Error
	if err != nil {
		return models.Connections{}, err
	}

	var connectionResponse models.ConnectionResponse
	json.Unmarshal(respBody, &connectionResponse)

	return connectionResponse.Connections, nil
}

// Get Connection - /api/v1/web_backend/connections/get
func GetConnection(connectionId string) (models.Connection, error) {
	var API_URL string = common.GetFullApiURL(GET_CONNECTION)

	respBody, err := common.ApiCall(API_URL, map[string]string{
		"connectionId": connectionId,
	})

	//Handle Error
	if err != nil {
		return models.Connection{}, err
	}

	var connectionResponse models.Connection
	json.Unmarshal(respBody, &connectionResponse)

	return connectionResponse, nil
}

func CreateConnection(connection models.ConnectionShort) (models.ConnectionShort, error) {
	var API_URL string = common.GetFullApiURL(CREATE_CONNECTION)

	var schedule interface{}

	schedule = connection.Schedule

	if connection.Manual {
		schedule = nil
	}

	respBody, err := common.ApiCallInterface(API_URL, map[interface{}]interface{}{
		"name":          connection.Name,
		"prefix":        connection.Prefix,
		"sourceId":      connection.SourceId,
		"destinationId": connection.DestinationId,
		"syncCatalog":   connection.Catalog,
		"schedule":      schedule,
		"status":        connection.Status,
	})

	var connectionResponse models.ConnectionShort
	json.Unmarshal(respBody, &connectionResponse)

	if err != nil {
		return models.ConnectionShort{}, err
	}

	return connectionResponse, nil
}
