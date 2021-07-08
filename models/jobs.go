package models

type Attempt struct {
	Id            int64
	Status        string
	CreatedAt     int
	UpdatedAt     int
	EndedAt       int
	BytesSynced   int
	RecordsSynced int
}

type Job struct {
	Job      JobDetail
	Attempts []Attempt
}

type JobDetail struct {
	Id         int64
	ConfigType string
	ConfigId   string
	CreatedAt  int
	UpdatedAt  int
	Status     string
}

type Jobs []Job

type JobsResponse struct {
	Jobs []Job
}

type ConfigTypes []string

type Pagination struct {
	PageSize  int `json:"pageSize"`
	RowOffset int `json:"rowOffset"`
}

type JobsPostBody struct {
	ConfigTypes []string   `json:configTypes`
	ConfigId    string     `json:configTypes`
	Pagination  Pagination `json:pagination`
}
