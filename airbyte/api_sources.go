package airbyte

import (
	"encoding/json"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/viper"
)

// Get Connections - /api/v1/sources/list
func GetSources() (models.Sources, error) {

	var API_URL string = common.GetFullApiURL(GET_SOURCES)

	workspaceId := viper.GetString("workspace_id")

	respBody, err := ApiCall(API_URL, map[string]string{
		"workspaceId": workspaceId,
	})

	//Handle Error
	if err != nil {
		return models.Sources{}, err
	}

	var sourceResponse models.SourceResponse
	json.Unmarshal(respBody, &sourceResponse)

	return sourceResponse.Sources, nil
}

// Get Connection - /api/v1/sources/get
func GetSource(sourceId string) (models.Source, error) {

	var API_URL string = common.GetFullApiURL(GET_SOURCE)

	respBody, err := ApiCall(API_URL, map[string]string{
		"sourceId": sourceId,
	})

	if err != nil {
		return models.Source{}, err
	}

	var sourceResponse models.Source
	json.Unmarshal(respBody, &sourceResponse)

	return sourceResponse, nil
}

// Check Source Connection - /api/v1/sources/check_connection
func CheckSourceConnection(sourceId string) (models.SourceCheckResponse, error) {
	var API_URL string = common.GetFullApiURL(SOURCE_CHECK_CONNECTION)

	respBody, err := ApiCall(API_URL, map[string]string{
		"sourceId": sourceId,
	})

	if err != nil {
		return models.SourceCheckResponse{}, err
	}

	var sourceCheckResponse models.SourceCheckResponse
	json.Unmarshal(respBody, &sourceCheckResponse)

	return sourceCheckResponse, nil
}
