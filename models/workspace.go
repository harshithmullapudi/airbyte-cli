package models

type Workspace struct {
	Name        string
	Email       string
	WorkspaceId string
	Slug        string
}

type Workspaces []Workspace

type WorkspaceResponse struct {
	Workspaces Workspaces
}
