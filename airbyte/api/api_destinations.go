package api

import (
	"encoding/json"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/viper"
)

func GetDestionations() (models.Destinations, error) {

	var API_URL string = common.GetFullApiURL(GET_DESTINATIONS)

	workspaceId := viper.GetString("workspace_id")

	respBody, err := common.ApiCall(API_URL, map[string]string{
		"workspaceId": workspaceId,
	})

	//Handle Error
	if err != nil {
		return models.Destinations{}, err
	}

	var destinationResponse models.DestinationResponse
	json.Unmarshal(respBody, &destinationResponse)

	return destinationResponse.Destinations, nil
}

func GetDestination(destinationId string) (models.Destination, error) {

	var API_URL string = common.GetFullApiURL(GET_DESTINATION)

	respBody, err := common.ApiCall(API_URL, map[string]string{
		"destinationId": destinationId,
	})

	//Handle Error
	if err != nil {
		return models.Destination{}, err
	}

	var destinationResponse models.Destination
	json.Unmarshal(respBody, &destinationResponse)

	return destinationResponse, nil
}
