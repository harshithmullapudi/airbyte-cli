package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

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

	defer resp.Body.Close()

	//Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
