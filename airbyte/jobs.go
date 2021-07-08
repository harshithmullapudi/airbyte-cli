package airbyte

import (
	"fmt"

	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
)

func PaginateJobs(configId string, configType string) (models.Jobs, error) {
	jobs, _ := GetJobs(configId, configType)
	logger.Log.Info("Fetching jobs from API for connection " + configId)
	logger.Log.Info("With Config Type: " + configType)
	logger.Log.Info("Total jobs: " + fmt.Sprintf("%d", len(jobs)))

	return jobs, nil
}
