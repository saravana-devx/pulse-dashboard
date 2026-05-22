package jobs

import (
	"errors"
)

var (
	ErrJobNotFound     = errors.New("job not found")
	ErrInvalidJobInput = errors.New("invalid job input")
	ErrInternal        = errors.New("internal error")
	ErrToCreateJob     = errors.New("failed to create job")
	ErrToUpdateJob     = errors.New("failed to update job")
	ErrToDeleteJob     = errors.New("failed to delete job")
	ErrUnauthorized    = errors.New("unauthorized")
)
