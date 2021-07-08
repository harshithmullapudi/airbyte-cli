package airbyte

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/models"
)

func GetJobs(configId string, configType string) (models.Jobs, error) {
	var API_URL string = common.GetFullApiURL(GET_JOBS)
	var configTypes []string

	configTypes = append(configTypes, configType)

	postBody, _ := json.Marshal(map[string]interface{}{
		"configTypes": configTypes,
		"configId":    configId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.Jobs{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Jobs{}, err
	}

	var jobsResponse models.JobsResponse
	json.Unmarshal(body, &jobsResponse)

	return jobsResponse.Jobs, nil
}

func GetJob(jobId int) (models.GetJobResponse, error) {
	var API_URL string = common.GetFullApiURL(GET_JOB)

	postBody, _ := json.Marshal(map[string]interface{}{
		"id": jobId,
	})

	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return models.GetJobResponse{}, err
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.GetJobResponse{}, err
	}

	var jobsResponse models.GetJobResponse
	json.Unmarshal(body, &jobsResponse)

	return jobsResponse, nil
}
