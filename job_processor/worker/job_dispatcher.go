package worker

import (
	"errors"
	"event-driven/common/model"
)

type JobDispatcher struct{}

func NewJobDispatcher() *JobDispatcher {
	return &JobDispatcher{}
}

func (d *JobDispatcher) Dispatch(jobType model.JobType, payload interface{}) (model.JobInterface, error) {
	job, exists := model.CreateJob(jobType, payload)
	if !exists {
		return nil, errors.New("unknown job type")
	}
	return job, nil
}

func initDispatcher() *JobDispatcher {
	return NewJobDispatcher()
}
