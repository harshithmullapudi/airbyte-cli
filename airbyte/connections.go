package airbyte

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
)

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

func FetchConnection(connectionId string) (models.Connection, error) {
	logger.Log.Info("Fetching connection from API for connectionId: " + connectionId)

	connection, err := GetConnection(connectionId)

	return connection, err
}

func SearchConnection(searchString string) (models.Connections, error) {
	if searchString == "" {
		return models.Connections{}, errors.New("you passed an empty string")
	}

	connections, _ := GetConnections()
	var filteredConnections models.Connections

	for _, c := range connections {
		if strings.Contains(strings.ToLower(c.Source.Name), strings.ToLower(searchString)) {
			filteredConnections = append(filteredConnections, c)
		}
	}

	return filteredConnections, nil
}
