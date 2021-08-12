package models

type ErrorResponse struct {
	Message string
}

type CheckResponse struct {
	Status  string
	Message string
	JobInfo JobInfo
}
