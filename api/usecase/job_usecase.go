package usecase

import (
	"event-driven/api/service"
	"event-driven/common/model"
)

type JobUsecase interface {
	EnqueueJob(job model.Job) (string, error)
	GetJobStatus(jobID string) (map[string]interface{}, error)
	UpdateJobStatus(jobID string, status string) error
}

type jobUsecase struct {
	jobService service.JobService
}

func NewJobUsecase(jobService service.JobService) JobUsecase {
	return &jobUsecase{jobService: jobService}
}

func (u *jobUsecase) EnqueueJob(job model.Job) (string, error) {
	return u.jobService.AddJobToQueue(job)
}

func (u *jobUsecase) GetJobStatus(jobID string) (map[string]interface{}, error) {
	return u.jobService.GetJobStatus(jobID)
}

func (u *jobUsecase) UpdateJobStatus(jobID string, status string) error {
	return u.jobService.UpdateJobStatus(jobID, status)
}
