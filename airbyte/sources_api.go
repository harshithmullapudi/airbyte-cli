package airbyte

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/spf13/viper"
)

func GetSources() (models.Sources, error) {

	var API_URL string = common.GetFullApiURL(GET_SOURCES)

	workspaceId := viper.GetString("workspace_id")

	postBody, _ := json.Marshal(map[string]string{
		"workspaceId": workspaceId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Sources{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Sources{}, err
	}

	var sourceResponse models.SourceResponse
	json.Unmarshal(body, &sourceResponse)

	return sourceResponse.Sources, nil
}

func GetSource(sourceId string) (models.Source, error) {

	var API_URL string = common.GetFullApiURL(GET_SOURCE)

	postBody, _ := json.Marshal(map[string]string{
		"sourceId": sourceId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Source{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Source{}, err
	}

	var sourceResponse models.Source
	json.Unmarshal(body, &sourceResponse)

	return sourceResponse, nil
}
