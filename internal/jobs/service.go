package jobs

import (
	"context"
)

type Service struct {
	repo *JobsRepository
}

func NewService(repo *JobsRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateJobService(ctx context.Context, req *CreateJobRequest) (*CreateJobResult, error) {
	return nil, nil
}

func (s *Service) GetJobByIdService(ctx context.Context, id string) (*GetJobByIdResult, error) {
	return nil, nil
}

func (s *Service) GetAllJobsService(ctx context.Context, userID string) (*GetAllJobsResult, error) {
	return nil, nil
}

func (s *Service) UpdateJobService(ctx context.Context, id string, req *UpdateJobRequest) (*Job, error) {
	return nil, nil
}

func (s *Service) DeleteJobService(ctx context.Context, id string) error {
	return nil
}
