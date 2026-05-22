package jobs

type CreateJobRequest struct{}

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
