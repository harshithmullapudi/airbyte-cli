package api

import (
	"encoding/json"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
)

func SourceConnectionCheckSchedule(source models.Source) (models.SourceCheckResponse, error) {
	var API_URL string = common.GetFullApiURL(SOURCE_CONNECTION_CHECK)
	respBody, err := common.ApiCallInterface(API_URL, map[interface{}]interface{}{
		"sourceDefinitionId":      source.SourceDefinitionId,
		"connectionConfiguration": source.ConnectionConfiguration,
	})

	var sourceCheckResponse models.SourceCheckResponse
	json.Unmarshal(respBody, &sourceCheckResponse)

	if err != nil {
		return models.SourceCheckResponse{}, err
	}

	return sourceCheckResponse, nil
}

func DestinationConnectionCheckSchedule(destination models.Destination) (models.CheckResponse, error) {
	var API_URL string = common.GetFullApiURL(DESTINATION_CONNECTION_CHECK)
	respBody, err := common.ApiCallInterface(API_URL, map[interface{}]interface{}{
		"destinationDefinitionId": destination.DestinationDefinitionId,
		"connectionConfiguration": destination.ConnectionConfiguration,
	})

	var destinationCheckResponse models.CheckResponse
	json.Unmarshal(respBody, &destinationCheckResponse)

	if err != nil {
		return models.CheckResponse{}, err
	}

	return destinationCheckResponse, nil
}
