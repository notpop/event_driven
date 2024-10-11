package model

import "log"

type JobType string

type Job struct {
	ID      string      `json:"id"`
	Type    JobType     `json:"type"`
	Payload interface{} `json:"payload"`
}

type JobInterface interface {
	Process() error
}

type jobEntry struct {
	jobType JobType
	creator func(payload interface{}) JobInterface
}

var jobRegistry = make(map[JobType]jobEntry)

func RegisterJob(jobType JobType, creator func(payload interface{}) JobInterface) {
	jobRegistry[jobType] = jobEntry{
		jobType: jobType,
		creator: creator,
	}
}

func CreateJob(jobType JobType, payload interface{}) (JobInterface, bool) {
	entry, exists := jobRegistry[jobType]
	if !exists {
		return nil, false
	}
	job := entry.creator(payload)
	return &wrappedJob{
		jobType: entry.jobType,
		job:     job,
	}, true
}

type wrappedJob struct {
	jobType JobType
	job     JobInterface
}

func (w *wrappedJob) Process() error {
	BeforeProcess(w.jobType)
	err := w.job.Process()
	AfterProcess(w.jobType)

	return err
}

func BeforeProcess(jobType JobType) {
	log.Printf("Starting job: Type=%s", jobType)
}

func AfterProcess(jobType JobType) {
	log.Printf("Finished job: Type=%s", jobType)
}
