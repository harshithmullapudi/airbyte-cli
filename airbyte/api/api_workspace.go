package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
)

func CheckIfWorkspaceExist(workspaceId string) (models.Workspace, error) {
	var API_URL string = common.GetFullApiURL(GET_WORKSPACES)

	postBody, _ := json.Marshal(map[string]string{
		"workspaceId": workspaceId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Workspace{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Workspace{}, err
	}

	var workspace models.Workspace
	json.Unmarshal(body, &workspace)

	return workspace, nil
}
