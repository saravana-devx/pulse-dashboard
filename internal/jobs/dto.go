package jobs

import (
	"time"

	"gorm.io/datatypes"
)

type CreateJobRequest struct {
	UserID      string         `json:"userId"`
	Type        JobType        `json:"type"`
	Payload     datatypes.JSON `json:"payload"`
	Priority    *int           `json:"priority,omitempty"`
	MaxRetries  *int           `json:"maxRetries,omitempty"`
	ScheduledAt *time.Time     `json:"scheduledAt,omitempty"`
}

type UpdateJobRequest struct{}

type GetJobByIdRequest struct{}

type GetAllJobsRequest struct{}

type DeleteJobRequest struct{}

type CreateJobResult struct {
	Job *Job
}

type GetJobByIdResult struct {
	Job *Job
}

type GetAllJobsResult struct {
	Jobs []*Job
}
