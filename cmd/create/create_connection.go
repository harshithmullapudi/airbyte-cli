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

var CONNECTIONS_CONFIG_FILE = "STANDARD_SYNC.yaml"

func CreateConnections(configFolderPath string, create bool) error {

	// Check if connections config file exists
	_, err := os.Stat(configFolderPath + "/" + CONNECTIONS_CONFIG_FILE)
	if os.IsNotExist(err) {
		logger.Warning("No config file found for connections. Skipping connections")
		return err
	}

	// Read config file
	var connections models.ConnectionsShort
	yamlFile, _ := ioutil.ReadFile(configFolderPath + "/" + CONNECTIONS_CONFIG_FILE)
	err = yaml.Unmarshal(yamlFile, &connections)
	if err != nil {
		logger.Errorf("Unmarshal: " + err.Error())
		return err
	}

	if !create {
		ValidateAllConnections(connections)
		return nil
	}

	for _, connection := range connections {
		err = CreateConnection(connection)

		if err != nil {
			logger.Error(err)
		}

	}

	return nil
}

func ValidateAllConnections(connections models.ConnectionsShort) error {
	for _, connection := range connections {
		err := ValidateConnection(connection.SourceId, connection.DestinationId)

		if err != nil {
			logger.Error(err)
		}

		logger.Notice("Connection " + connection.SourceId + " - " + connection.DestinationId + " is valid")
	}

	return nil
}

func ValidateConnection(sourceId string, destinationId string) error {
	// Check if source exists
	_, err := api.GetSource(sourceId)

	if err != nil {
		return errors.New("Source not found for connection " + sourceId)
	}

	// Check if destination exists
	_, err = api.GetDestination(destinationId)

	if err != nil {
		return errors.New("Destination not found for connection " + destinationId)
	}

	return nil
}

func CreateConnection(connection models.ConnectionShort) error {

	// Validate i
	err := ValidateConnection(connection.SourceId, connection.DestinationId)

	if err != nil {
		return err
	}

	if connection.ConnectionId != "" {
		_, err := api.GetConnection(connection.ConnectionId)

		if err == nil {
			return errors.New("connection already exist")
		}
	}

	_, err = api.CreateConnection(connection)

	if err != nil {
		return errors.New("creation for connection failed " + err.Error())
	}

	logger.Notice("Connection " + connection.SourceId + " - " + connection.DestinationId + " created successfully")

	return nil
}
