package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
)

// Get Workspaces - /v1/workspaces/list
func GetWorkspaces() (models.Workspaces, error) {

	var API_URL string = common.GetFullApiURL(GET_WORKSPACES)

	respBody, err := common.ApiCall(API_URL, map[string]string{})

	//Handle Error
	if err != nil {
		return models.Workspaces{}, err
	}

	var workspaceResponse models.WorkspaceResponse
	json.Unmarshal(respBody, &workspaceResponse)

	return workspaceResponse.Workspaces, nil
}

// Get Connection - /api/v1/workspaces/get
func GetWorkspace(workspaceId string) (models.Workspace, error) {

	//Get First Workspace
	if workspaceId == "" {
		workspaces, err := GetWorkspaces()

		//Handle Error
		if err != nil {
			return models.Workspace{}, err
		}

		workspaceId = workspaces[0].WorkspaceId
	}

	var API_URL string = common.GetFullApiURL(GET_WORKSPACE)

	respBody, err := common.ApiCall(API_URL, map[string]string{
		"workspaceId": workspaceId,
	})

	if err != nil {
		return models.Workspace{}, err
	}

	var workspaceResponse models.Workspace
	json.Unmarshal(respBody, &workspaceResponse)

	return workspaceResponse, nil
}

func CheckIfWorkspaceExist(workspaceId string) (models.Workspace, error) {
	var API_URL string = common.GetFullApiURL(GET_WORKSPACE)

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
