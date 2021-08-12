package create

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/harshithmullapudi/airbyte/airbyte/api"
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
	"gopkg.in/yaml.v2"
)

var SOURCES_CONFIG_FILE = "SOURCE_CONNECTION.yaml"

func CreateSources(configFolderPath string, create bool) error {

	// Check if sources config file exists
	_, err := os.Stat(configFolderPath + "/" + SOURCES_CONFIG_FILE)
	if os.IsNotExist(err) {
		logger.Errorf("No config file found for sources. Skipping sources")
		return err
	}

	// Read config file
	var sources models.Sources
	yamlFile, _ := ioutil.ReadFile(configFolderPath + "/" + SOURCES_CONFIG_FILE)
	err = yaml.Unmarshal(yamlFile, &sources)
	if err != nil {
		logger.Errorf("Unmarshal: " + err.Error())
		return err
	}

	if !create {
		ValidateAllSources(sources)
		return nil
	}

	for _, source := range sources {
		err = CreateSource(source)

		if err != nil {
			logger.Error("Source creation for " + source.Name + " has failed. " + err.Error() + " . Skipping creation for this source")
		}

	}

	return nil
}

func ValidateAllSources(sources models.Sources) error {
	for _, source := range sources {
		sourceCheckResponse, err := api.SourceConnectionCheckSchedule(source)

		if err != nil {
			logger.Error("Source " + source.Name + " has failed during validation with " + err.Error())
			continue
		}

		if sourceCheckResponse.Status != "succeeded" {
			logger.Error("Source " + source.Name + " has failed during validation with " + sourceCheckResponse.Message)
			continue
		}

		logger.Notice("Source " + source.Name + " has passed validation")
	}

	return nil
}

func CreateSource(source models.Source) error {

	// Validate i
	sourceCheckResponse, err := api.SourceConnectionCheckSchedule(source)

	if err != nil {
		return err
	}

	if sourceCheckResponse.Status != "succeeded" {
		return errors.New("validation for source failed")
	}

	if source.SourceId != "" {
		_, err := api.GetSource(source.SourceId)

		if err == nil {
			return errors.New("source already exist")
		}
	}

	_, err = api.CreateSource(source)

	if err != nil {
		return errors.New("creation for source failed")
	}

	logger.Notice("Source " + source.Name + " created successfully")

	return nil
}
