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

func GetDestionations() (models.Destinations, error) {

	var API_URL string = common.GetFullApiURL(GET_DESTINATIONS)

	workspaceId := viper.GetString("workspace_id")

	postBody, _ := json.Marshal(map[string]string{
		"workspaceId": workspaceId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Destinations{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Destinations{}, err
	}

	var destinationResponse models.DestinationResponse
	json.Unmarshal(body, &destinationResponse)

	return destinationResponse.Destinations, nil
}

func GetDestination(destinationId string) (models.Destination, error) {

	var API_URL string = common.GetFullApiURL(GET_DESTINATION)

	postBody, _ := json.Marshal(map[string]string{
		"destinationId": destinationId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Destination{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Destination{}, err
	}

	var destinationResponse models.Destination
	json.Unmarshal(body, &destinationResponse)

	return destinationResponse, nil
}
