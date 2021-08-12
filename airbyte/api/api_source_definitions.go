package api

import (
	"encoding/json"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
)

func GetSourceDefinitions() (models.SourceDefinitions, error) {
	var API_URL string = common.GetFullApiURL(GET_SOURCE_DEFINITIONS)

	respBody, err := common.ApiCall(API_URL, nil)

	//Handle Error
	if err != nil {
		return models.SourceDefinitions{}, err
	}

	var sourceDefinitionResponse models.SourceDefinitionResponse
	json.Unmarshal(respBody, &sourceDefinitionResponse)

	return sourceDefinitionResponse.SourceDefinitions, nil
}
