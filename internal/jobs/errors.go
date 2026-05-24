package jobs

import (
	"errors"
)

var (
	ErrJobNotFound     = errors.New("job not found")
	ErrInvalidJobInput = errors.New("invalid job input")
	ErrInternal        = errors.New("internal error")
	ErrToCreateJob     = errors.New("failed to create job")
	ErrToCreateJobs    = errors.New("failed to create jobs")
	ErrToGetJob        = errors.New("failed to get job")
	ErrToGetAllJobs    = errors.New("failed to get jobs")
	ErrToUpdateJob     = errors.New("failed to update job")
	ErrToDeleteJob     = errors.New("failed to delete job")
	ErrUnauthorized    = errors.New("unauthorized")
)
