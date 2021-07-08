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

func GetConnections() (models.Connections, error) {
	var API_URL string = common.GetFullApiURL(GET_CONNECTIONS)

	workspaceId := viper.GetString("workspace_id")

	postBody, _ := json.Marshal(map[string]string{
		"workspaceId": workspaceId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Connections{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Connections{}, err
	}

	var connectionResponse models.ConnectionResponse
	json.Unmarshal(body, &connectionResponse)

	return connectionResponse.Connections, nil
}

func GetConnection(connectionId string) (models.Connection, error) {
	var API_URL string = common.GetFullApiURL(GET_CONNECTION)

	postBody, _ := json.Marshal(map[string]string{
		"connectionId": connectionId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Connection{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Connection{}, err
	}

	var connectionResponse models.Connection
	json.Unmarshal(body, &connectionResponse)

	return connectionResponse, nil
}
