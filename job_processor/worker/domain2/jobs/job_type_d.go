package jobs

import (
	"event-driven/common/model"
	"log"
)

type JobTypeD struct {
	Payload map[string]interface{}
}

func (j *JobTypeD) Process() error {
	log.Printf("Processing Job Type D with payload: %v", j.Payload)
	return nil
}

func init() {
	model.RegisterJob("Domain2JobTypeD", func(payload interface{}) model.JobInterface {
		return &JobTypeD{Payload: payload.(map[string]interface{})}
	})
}
