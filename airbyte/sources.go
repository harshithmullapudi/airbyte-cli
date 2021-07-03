package airbyte

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/harshithmullapudi/airbyte/common"
	"github.com/harshithmullapudi/airbyte/logger"
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

func PaginateSources(offset int, number int) (models.Sources, error) {
	sources, _ := GetSources()
	logger.Log.Info("Fetching sources from API with offset: " + fmt.Sprintf("%d", offset))
	logger.Log.Info("Total sources: " + fmt.Sprintf("%d", len(sources)))

	var endIndex int = int(math.Min(float64(offset+number), float64(len(sources))))

	var finalSources models.Sources = sources[offset:endIndex]
	return finalSources, nil
}

func GetSource(sourceId string) (models.Source, error) {
	logger.Log.Info("Fetching source from API for sourceId: " + sourceId)

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

func SearchSource(searchString string) (models.Sources, error) {
	if searchString == "" {
		return models.Sources{}, errors.New("you passed an empty string")
	}

	sources, _ := GetSources()
	var filteredSources models.Sources

	for _, s := range sources {
		if strings.Contains(strings.ToLower(s.Name), strings.ToLower(searchString)) {
			filteredSources = append(filteredSources, s)
		}
	}

	return filteredSources, nil
}
