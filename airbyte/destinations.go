package airbyte

import (
	"fmt"
	"math"

	"github.com/harshithmullapudi/airbyte/airbyte/api"
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
)

func PaginateDestinations(offset int, number int) (models.Destinations, error) {
	destinations, _ := api.GetDestionations()
	logger.Log.Info("Fetching sources from API with offset: " + fmt.Sprintf("%d", offset))
	logger.Log.Info("Total sources: " + fmt.Sprintf("%d", len(destinations)))

	var endIndex int = int(math.Min(float64(offset+number), float64(len(destinations))))

	var finalDestinations models.Destinations = destinations[offset:endIndex]
	return finalDestinations, nil
}

func FetchDestination(destinationId string) (models.Destination, error) {
	logger.Log.Info("Fetching destination from API for destinationId: " + destinationId)

	destination, err := api.GetDestination(destinationId)
	return destination, err
}
