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

func PaginateWorkspaces(offset int, number int) (models.Workspaces, error) {
	workspaces, _ := api.GetWorkspaces()
	logger.Log.Info("Fetching workspaces from API with offset: " + fmt.Sprintf("%d", offset))
	logger.Log.Info("Total workspaces: " + fmt.Sprintf("%d", len(workspaces)))

	var endIndex int = int(math.Min(float64(offset+number), float64(len(workspaces))))

	var finalWorkspaces models.Workspaces = workspaces[offset:endIndex]
	return finalWorkspaces, nil
}

func FetchWorkspace(workspaceId string) (models.Workspace, error) {
	logger.Log.Info("Fetching workspace from API for workspaceId: " + workspaceId)

	workspace, err := api.GetWorkspace(workspaceId)
	return workspace, err
}

func SearchWorkspace(searchString string) (models.Workspaces, error) {
	if searchString == "" {
		return models.Workspaces{}, errors.New("you passed an empty string")
	}

	workspaces, _ := api.GetWorkspaces()
	var filteredWorkspaces models.Workspaces

	for _, s := range workspaces {
		if strings.Contains(strings.ToLower(s.Name), strings.ToLower(searchString)) {
			filteredWorkspaces = append(filteredWorkspaces, s)
		}
	}

	return filteredWorkspaces, nil
}
