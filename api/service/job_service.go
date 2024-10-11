package service

import (
	"event-driven/common/model"
	"event-driven/queue"

	"github.com/google/uuid"
)

type JobService interface {
	AddJobToQueue(job model.Job) (string, error)
	GetJobStatus(jobID string) (map[string]interface{}, error)
	UpdateJobStatus(jobID string, status string) error
}

type jobService struct{}

func NewJobService() JobService {
	return &jobService{}
}

func (s *jobService) AddJobToQueue(job model.Job) (string, error) {
	jobID := uuid.New().String()
	job.ID = jobID
	err := queue.AddJobToQueue(job)
	return jobID, err
}

func (s *jobService) GetJobStatus(jobID string) (map[string]interface{}, error) {
	return queue.GetJobStatus(jobID)
}

func (s *jobService) UpdateJobStatus(jobID string, status string) error {
	return queue.UpdateJobStatus(jobID, status)
}
