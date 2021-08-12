package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/harshithmullapudi/airbyte/models"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

func GetFullApiURL(path string) string {
	airbyte_url := viper.GetString("airbyte_url")
	return airbyte_url + path
}

func ApiCall(API_URL string, jsonBody map[string]string) ([]byte, error) {

	var body []byte

	if len(jsonBody) != 0 {
		body, _ = json.Marshal(jsonBody)

	} else {
		body, _ = json.Marshal(map[string]string{})
	}

	requestBody := bytes.NewBuffer(body)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		//Read the response body
		errBody, _ := ioutil.ReadAll(resp.Body)
		var errorResponse models.ErrorResponse
		json.Unmarshal(errBody, &errorResponse)
		return nil, errors.New(errorResponse.Message)
	}

	defer resp.Body.Close()

	//Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func ApiCallInterface(API_URL string, jsonBody map[interface{}]interface{}) ([]byte, error) {

	var body []byte
	var err error

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if len(jsonBody) != 0 {
		body, _ = json.Marshal(jsonBody)
	} else {
		body, _ = json.Marshal(map[string]string{})
	}

	requestBody := bytes.NewBuffer(body)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(API_URL, "application/json", requestBody)

	//Handle Error
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		//Read the response body
		errBody, _ := ioutil.ReadAll(resp.Body)
		var errorResponse models.ErrorResponse
		json.Unmarshal(errBody, &errorResponse)
		return nil, errors.New(errorResponse.Message)
	}

	defer resp.Body.Close()

	//Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
