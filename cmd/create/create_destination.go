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

var DESTINATIONS_CONFIG_FILE = "DESTINATION_CONNECTION.yaml"

func CreateDestinations(configFolderPath string, create bool) error {

	// Check if destinations config file exists
	_, err := os.Stat(configFolderPath + "/" + DESTINATIONS_CONFIG_FILE)
	if os.IsNotExist(err) {
		logger.Warning("No config file found for destinations. Skipping destinations")
		return err
	}

	// Read config file
	var destinations models.Destinations
	yamlFile, _ := ioutil.ReadFile(configFolderPath + "/" + DESTINATIONS_CONFIG_FILE)
	err = yaml.Unmarshal(yamlFile, &destinations)
	if err != nil {
		logger.Errorf("Unmarshal: " + err.Error())
		return err
	}

	if !create {
		ValidateAllDestinations(destinations)
		return nil
	}

	for _, destination := range destinations {
		err = CreateDestination(destination)

		if err != nil {
			logger.Error("Destination creation for " + destination.Name + " has failed. " + err.Error() + ". Skipping creation for this destination")
		}

	}

	return nil
}

func ValidateAllDestinations(destinations models.Destinations) error {
	for _, destination := range destinations {
		destinationCheckResponse, err := api.DestinationConnectionCheckSchedule(destination)

		if err != nil {
			logger.Error("Destination " + destination.Name + " has failed during validation with " + err.Error())
			continue
		}

		if destinationCheckResponse.Status != "succeeded" {
			logger.Error("Destination " + destination.Name + " has failed during validation with " + destinationCheckResponse.Message)
			continue
		}

		logger.Notice("Destination " + destination.Name + " has passed validation")
	}

	return nil
}

func CreateDestination(destination models.Destination) error {

	// Validate i
	destinationCheckResponse, err := api.DestinationConnectionCheckSchedule(destination)

	if err != nil {
		return err
	}

	if destinationCheckResponse.Status != "succeeded" {
		return errors.New("validation for destination failed")
	}

	if destination.DestinationId != "" {
		_, err := api.GetDestination(destination.DestinationId)

		if err == nil {
			return errors.New("destination already exist")
		}
	}

	_, err = api.CreateDestination(destination)

	if err != nil {
		return errors.New("creation for destination failed")
	}

	logger.Notice("Destination " + destination.Name + " created successfully")

	return nil
}
