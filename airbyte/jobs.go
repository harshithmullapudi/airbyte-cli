package airbyte

import (
	"fmt"

	"github.com/harshithmullapudi/airbyte/airbyte/api"
	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
)

func PaginateJobs(configId string, configType string) (models.Jobs, error) {
	logger.Log.Info("Fetching jobs from API for connection " + configId)
	logger.Log.Info("With Config Type: " + configType)
	jobs, _ := api.GetJobs(configId, configType)
	logger.Log.Info("Total jobs: " + fmt.Sprintf("%d", len(jobs)))

	return jobs, nil
}

func FetchJob(jobId int) (models.GetJobResponse, error) {
	logger.Log.Info("Fetching job from API for id " + fmt.Sprint(jobId))
	job, _ := api.GetJob(jobId)

	return job, nil
}
