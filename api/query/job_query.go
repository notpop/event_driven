package query

import (
	"event-driven/common/model"
	"event-driven/queue"
)

type JobQuery interface {
	GetJobs() ([]model.Job, error)
}

type jobQuery struct{}

func NewJobQuery() JobQuery {
	return &jobQuery{}
}

func (q *jobQuery) GetJobs() ([]model.Job, error) {
	return queue.GetAllJobs()
}
