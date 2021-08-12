package airbyte

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/harshithmullapudi/airbyte/airbyte/api"
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
)

func PaginateSources(offset int, number int) (models.Sources, error) {
	sources, _ := api.GetSources()
	logger.Log.Info("Fetching sources from API with offset: " + fmt.Sprintf("%d", offset))
	logger.Log.Info("Total sources: " + fmt.Sprintf("%d", len(sources)))

	var endIndex int = int(math.Min(float64(offset+number), float64(len(sources))))

	var finalSources models.Sources = sources[offset:endIndex]
	return finalSources, nil
}

func FetchSource(sourceId string) (models.Source, error) {
	logger.Log.Info("Fetching source from API for sourceId: " + sourceId)

	source, err := api.GetSource(sourceId)
	return source, err
}

func SearchSource(searchString string) (models.Sources, error) {
	if searchString == "" {
		return models.Sources{}, errors.New("you passed an empty string")
	}

	sources, _ := api.GetSources()
	var filteredSources models.Sources

	for _, s := range sources {
		if strings.Contains(strings.ToLower(s.Name), strings.ToLower(searchString)) {
			filteredSources = append(filteredSources, s)
		}
	}

	return filteredSources, nil
}
