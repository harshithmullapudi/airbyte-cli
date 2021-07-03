package airbyte

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/logger"
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

func PaginateConnections(offset int, number int, status string) (models.Connections, error) {
	connections, _ := GetConnections()
	logger.Log.Info("Fetching connections from API with offset: " + fmt.Sprintf("%d", offset))
	logger.Log.Info("Total connections: " + fmt.Sprintf("%d", len(connections)))

	var filteredConnections models.Connections

	// filter status attribute
	if status == "" {
		filteredConnections = connections
	} else {
		for _, c := range connections {
			if c.Status == status {
				filteredConnections = append(filteredConnections, c)
			}
		}
	}

	var endIndex int = int(math.Min(float64(offset+number), float64(len(filteredConnections))))

	var finalConnections models.Connections = filteredConnections[offset:endIndex]
	return finalConnections, nil
}

func GetConnection(connectionId string) (models.Connection, error) {
	logger.Log.Info("Fetching connection from API for connectionId: " + connectionId)

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
