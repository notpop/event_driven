package jobs

import (
	"event-driven/common/model"
	"log"
)

type JobTypeC struct {
	Payload map[string]interface{}
}

func (j *JobTypeC) Process() error {
	log.Printf("Processing Job Type D with payload: %v", j.Payload)
	return nil
}

func init() {
	model.RegisterJob("Domain2JobTypeC", func(payload interface{}) model.JobInterface {
		return &JobTypeC{Payload: payload.(map[string]interface{})}
	})
}
